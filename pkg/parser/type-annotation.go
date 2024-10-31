package parser

import (
	"reflect"

	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseTypeAnnotation(v ASTNode) string {
	result := ""
	vType := reflect.TypeOf(v)

	switch vType {
	case reflect.TypeOf(VariableDeclaration{}):
		node := v.(VariableDeclaration)
		_, valueType := p.ReflectInitVariableDeclaratorType(node)

		if node.TypeAnnotation != "" {
			if valueType != node.TypeAnnotation {
				oops.TypeMismatch(p.peek(), node.TypeAnnotation, valueType)
			}

			result = node.TypeAnnotation
		} else {
			result = p.InferType(v)
		}
	default:
		break
	}

	return result
}

func (p *Parser) InferType(v ASTNode) string {
	result := ""
	vType := reflect.TypeOf(v)

	switch vType {
	case reflect.TypeOf(VariableDeclaration{}):
		node := v.(VariableDeclaration)
		_, valueType := p.ReflectInitVariableDeclaratorType(node)
		result = valueType
	default:
		result = ""
	}

	return result
}

func (p *Parser) ReflectInitVariableDeclaratorType(v VariableDeclaration) (reflect.Type, string) {
	initType := reflect.TypeOf(v.Declarations[0].Init)
	valueType := ""

	if initType.Kind() == reflect.Struct {
		switch initType {
		case reflect.TypeOf(Literal{}):
			valueType = reflect.TypeOf(v.Declarations[0].Init.(Literal).Value).String()
		case reflect.TypeOf(BinaryExpression{}):
			isConcat := p.IsConcatenation(v.Declarations[0].Init.(BinaryExpression))

			if isConcat {
				valueType = "string"
				break
			}

			result := p.EvaluateBinaryExpr(v.Declarations[0].Init.(BinaryExpression))
			valueType = reflect.TypeOf(result).String()
			break
		default:
			valueType = ""
		}
	}

	return initType, valueType
}

func (p *Parser) IsConcatenation(b ASTNode) bool {
	if reflect.TypeOf(b) == reflect.TypeOf(Literal{}) {
		return reflect.TypeOf(b.(Literal).Value) == reflect.TypeOf("")
	}

	if reflect.TypeOf(b) == reflect.TypeOf(BinaryExpression{}) {
		leftTypeIsString := p.IsConcatenation(b.(BinaryExpression).Left)
		rightTypeIsString := p.IsConcatenation(b.(BinaryExpression).Right)
		isPlusSign := b.(BinaryExpression).Operator == "+"

		if (leftTypeIsString || rightTypeIsString) && !isPlusSign {
			oops.InvalidConcatenation(p.peek(), ""+b.(BinaryExpression).Operator+" operator"+" is not allowed")
		}

		return (leftTypeIsString || rightTypeIsString) && isPlusSign
	}

	return false
}

func (p *Parser) EvaluateBinaryExpr(b ASTNode) interface{} {
	if reflect.TypeOf(b) == reflect.TypeOf(Literal{}) {
		return b.(Literal).Value
	}

	if reflect.TypeOf(b) == reflect.TypeOf(BinaryExpression{}) {
		leftValue := p.EvaluateBinaryExpr(b.(BinaryExpression).Left)
		rightValue := p.EvaluateBinaryExpr(b.(BinaryExpression).Right)
		operator := b.(BinaryExpression).Operator

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
