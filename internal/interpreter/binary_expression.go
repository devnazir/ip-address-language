package interpreter

import (
	"fmt"
	"strings"

	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
	"github.com/devnazir/ip-address-language/pkg/utils"
)

func (i *Interpreter) InterpretBinaryExpr(b ast.ASTNode, returnAsRaw bool) interface{} {
	switch b.GetType() {
	case ast.StringLiteralTree:
		if returnAsRaw {
			return b.(ast.StringLiteral).Value
		}

		return b.(ast.StringLiteral).Value

	case ast.NumberLiteralTree:
		return b.(ast.NumberLiteral).Value
	case ast.BooleanLiteralTree:
		return b.(ast.BooleanLiteral).Value
	case ast.IdentifierTree:
		name := b.(ast.Identifier).Name
		info := i.scopeResolver.ResolveScope(name)
		value := fmt.Sprintf("%v", info.Value)
		finalValue := strings.NewReplacer(name, value).Replace(b.(ast.Identifier).Raw)

		if returnAsRaw {
			return finalValue
		}

		return value

	case ast.SubShellTree:
		value := i.InterpretSubShell(b.(ast.SubShell).Arguments)
		return value

	case ast.MemberExpressionTree:
		memberExpr := b.(ast.MemberExpression)
		value := i.EvaluateMemberExpr(memberExpr, memberExpr.Computed)
		return value

	case ast.BinaryExpressionTree:
		leftValue := i.InterpretBinaryExpr(b.(ast.BinaryExpression).Left, returnAsRaw)
		rightValue := i.InterpretBinaryExpr(b.(ast.BinaryExpression).Right, returnAsRaw)
		operator := b.(ast.BinaryExpression).Operator

		leftValue, leftType := utils.InferType(leftValue)
		rightValue, rightType := utils.InferType(rightValue)

		i.validateOperandTypes(leftType, rightType, operator)

		if leftType == lx.StringType && rightType == lx.StringType {
			leftStr := leftValue.(string)
			rightStr := rightValue.(string)

			switch operator {
			case lx.EquivalenceSign:
				return leftStr == rightStr
			case lx.NotEqualsSign:
				return leftStr != rightStr
			case lx.GreaterThanSign:
				return leftStr > rightStr
			case lx.LessThanSign:
				return leftStr < rightStr
			case lx.GreaterOrEqualSign:
				return leftStr >= rightStr
			case lx.LessOrEqualSign:
				return leftStr <= rightStr
			case lx.AdditionSign:
				return i.ConcatenateString(b)
			default:
				panic(oops.SyntaxError(b.(ast.BinaryExpression).Right, "Invalid operation"))
			}
		}

		if leftType == lx.BoolType && rightType == lx.BoolType {
			leftBool := leftValue.(bool)
			rightBool := rightValue.(bool)

			switch operator {
			case lx.AndOperator:
				return leftBool && rightBool
			case lx.OrOperator:
				return leftBool || rightBool
			default:
				panic(oops.SyntaxError(b.(ast.BinaryExpression).Right, "Invalid operation"))
			}
		}

		leftFloatValue := i.toFloat64(leftValue, leftType)
		rightFloatValue := i.toFloat64(rightValue, rightType)
		isLeftOrRightFloat := leftType == lx.Float64Type || rightType == lx.Float64Type

		switch operator {
		case lx.EquivalenceSign:
			return leftValue == rightValue
		case lx.NotEqualsSign:
			return leftValue != rightValue
		case lx.GreaterThanSign:
			return i.compareValues(leftValue, rightValue, leftType, rightType, func(l, r float64) bool { return l > r })
		case lx.LessThanSign:
			return i.compareValues(leftValue, rightValue, leftType, rightType, func(l, r float64) bool { return l < r })
		case lx.GreaterOrEqualSign:
			return i.compareValues(leftValue, rightValue, leftType, rightType, func(l, r float64) bool { return l >= r })
		case lx.LessOrEqualSign:
			return i.compareValues(leftValue, rightValue, leftType, rightType, func(l, r float64) bool { return l <= r })
		case lx.AndOperator:

			if leftType != lx.BoolType || rightType != lx.BoolType {
				panic(oops.SyntaxError(b.(ast.BinaryExpression).Right, "Invalid operation: both operands must be booleans for boolean comparison"))
			}

			return leftValue.(bool) && rightValue.(bool)
		case lx.OrOperator:

			if leftType != lx.BoolType || rightType != lx.BoolType {
				panic(oops.SyntaxError(b.(ast.BinaryExpression).Right, "Invalid operation: both operands must be booleans for boolean comparison"))
			}

			return leftValue.(bool) || rightValue.(bool)
		case lx.AdditionSign:
			if isLeftOrRightFloat {
				return leftFloatValue + rightFloatValue
			}

			return int(leftFloatValue + rightFloatValue)

		case lx.SubtractionSign:
			if isLeftOrRightFloat {
				return leftFloatValue - rightFloatValue
			}

			return int(leftFloatValue - rightFloatValue)

		case lx.MultiplicationSign:
			if isLeftOrRightFloat {
				return leftFloatValue * rightFloatValue
			}

			return int(leftFloatValue * rightFloatValue)

		case lx.DivisionSign:
			if isLeftOrRightFloat {
				return leftFloatValue / rightFloatValue
			}

			return int(leftFloatValue / rightFloatValue)

		case lx.ModulusSign:
			return int(leftFloatValue) % int(rightFloatValue)

		default:
			panic(oops.SyntaxError(b.(ast.BinaryExpression).Right, "Invalid operation"))
		}
	}

	return nil
}

func (i *Interpreter) ConcatenateString(b ast.ASTNode) string {
	switch b.GetType() {
	case ast.StringLiteralTree:
		return b.(ast.StringLiteral).Value

	case ast.IdentifierTree:
		return strings.TrimSpace(i.resolveIdentifier(b.(ast.Identifier)).(string))
	case ast.BinaryExpressionTree:
		leftValue := i.ConcatenateString(b.(ast.BinaryExpression).Left)
		rightValue := i.ConcatenateString(b.(ast.BinaryExpression).Right)

		return leftValue + rightValue
	default:
		return ""
	}
}

func (i *Interpreter) validateOperandTypes(leftType, rightType, operator string) {
	if operator == lx.EquivalenceSign || operator == lx.NotEqualsSign || operator == lx.GreaterThanSign || operator == lx.LessThanSign || operator == lx.GreaterOrEqualSign || operator == lx.LessOrEqualSign || operator == lx.AdditionSign {
		if leftType == lx.StringType || rightType == lx.StringType {
			if leftType != lx.StringType || rightType != lx.StringType {
				panic("Invalid operation: both operands must be strings for string comparison")
			}
		}
	}

	if leftType == lx.BoolType || rightType == lx.BoolType {
		if leftType != lx.BoolType || rightType != lx.BoolType {
			panic("Invalid operation: both operands must be booleans for boolean comparison")
		}
	}
}

func (i *Interpreter) compareValues(leftValue, rightValue interface{}, leftType, rightType string, comparator func(float64, float64) bool) bool {
	left := i.toFloat64(leftValue, leftType)
	right := i.toFloat64(rightValue, rightType)
	return comparator(left, right)
}

func (i *Interpreter) toFloat64(value interface{}, valueType string) float64 {
	switch valueType {
	case lx.IntType:
		return float64(value.(int))
	case lx.Float64Type:
		return value.(float64)
	default:
		panic("Unsupported type for numeric comparison")
	}
}
