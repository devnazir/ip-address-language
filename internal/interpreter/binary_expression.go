package interpreter

import (
	"reflect"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (i *Interpreter) InterpretBinaryExpr(b ast.ASTNode) interface{} {
	if reflect.TypeOf(b) == reflect.TypeOf(ast.Literal{}) {
		return b.(ast.Literal).Value
	}

	if reflect.TypeOf(b) == reflect.TypeOf(ast.Identifier{}) {
		name := b.(ast.Identifier).Name
		value := env.GetVariable(name)

		if value == nil {
			oops.IdentifierNotFoundError(b.(ast.Identifier))
		}

		return value
	}

	if reflect.TypeOf(b) == reflect.TypeOf(ast.BinaryExpression{}) {
		leftValue := i.InterpretBinaryExpr(b.(ast.BinaryExpression).Left)
		rightValue := i.InterpretBinaryExpr(b.(ast.BinaryExpression).Right)
		operator := b.(ast.BinaryExpression).Operator

		var leftFloat, rightFloat float64
		var isLeftFloat, isRightFloat bool

		switch v := leftValue.(type) {
		case int:
			leftFloat = float64(v)
		case float64:
			leftFloat = v
			isLeftFloat = true
		default:
			return 0
		}

		switch v := rightValue.(type) {
		case int:
			rightFloat = float64(v)
		case float64:
			rightFloat = v
			isRightFloat = true
		default:
			return 0
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
	}

	return 0
}
