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

			// TODO: Implement binary expression type inference
			valueType = "int"

		default:
			valueType = ""
		}
	}

	return initType, valueType
}
