package interpreter

import (
	"fmt"
	"reflect"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (i *Interpreter) InterpretBinaryExpr(b ast.ASTNode) interface{} {

	switch b.(type) {
	case ast.StringLiteral:
		return b.(ast.StringLiteral).Value
	case ast.NumberLiteral:
		return b.(ast.NumberLiteral).Value
	case ast.Identifier:
		name := b.(ast.Identifier).Name
		info := i.scopeResolver.ResolveScope(name)

		return info.Value

	case ast.SubShell:
		value := i.InterpretSubShell(b.(ast.SubShell).Arguments.(string))
		return value

	case ast.MemberExpression:
		memberExpr := b.(ast.MemberExpression)
		value := i.EvaluateMemberExpr(memberExpr, memberExpr.Computed)
		return value

	case ast.BinaryExpression:
		leftValue := i.InterpretBinaryExpr(b.(ast.BinaryExpression).Left)
		rightValue := i.InterpretBinaryExpr(b.(ast.BinaryExpression).Right)
		operator := b.(ast.BinaryExpression).Operator
		isConcat := false

		var leftFloat, rightFloat float64
		var isLeftFloat, isRightFloat bool

		if reflect.TypeOf(leftValue) == reflect.TypeOf("") || reflect.TypeOf(rightValue) == reflect.TypeOf("") {
			isConcat = true
		}

		switch v := leftValue.(type) {
		case int:
			leftFloat = float64(v)
		case float64:
			leftFloat = v
			isLeftFloat = true
		case string:
			leftValue = string(v)
		default:
			return v
		}

		switch v := rightValue.(type) {
		case int:
			rightFloat = float64(v)
		case float64:
			rightFloat = v
			isRightFloat = true
		case string:
			rightValue = string(v)
		default:
			return v
		}

		if isConcat {

			if operator != "+" {
				panic(fmt.Sprintf("%v operator is not allowed", operator))
			}

			return fmt.Sprintf("%v", leftValue) + fmt.Sprintf("%v", rightValue)
		}

		var result interface{}
		switch operator {
		case "+":
			result = leftFloat + rightFloat
		case "-":
			result = leftFloat - rightFloat
		case "*":
			result = leftFloat * rightFloat
		case "/":
			if rightFloat == 0 {
				return 0
			}
			result = leftFloat / rightFloat
		}

		if isLeftFloat || isRightFloat {
			return result
		}
		return int(result.(float64))

	default:
		return 0
	}
}
