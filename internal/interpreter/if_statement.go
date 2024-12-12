package interpreter

import (
	"reflect"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func isBooleanValue(value interface{}) bool {
	_, ok := value.(bool)
	return ok
}

func (i *Interpreter) InterpretIfStatement(node ast.IfStatement) {
	condition := i.InterpretBinaryExpr(node.Condition, false)
	consequent := node.Consequent

	if condition == nil {
		panic(oops.RuntimeError(node.Condition, "Invalid condition"))
	}

	if !isBooleanValue(condition) {
		panic(oops.RuntimeError(node.Condition, "Condition must be a boolean value"))
	}

	if condition.(bool) {
		if _, ok := consequent.(ast.BodyProgram); ok {
			for _, statement := range consequent.(ast.BodyProgram) {
				i.InterpretNode(statement, "")
			}
		}
	} else {
		alternate := node.Alternate

		if reflect.TypeOf(alternate) == reflect.TypeOf(ast.IfStatement{}) {
			i.InterpretIfStatement(alternate.(ast.IfStatement))
		} else {
			if _, ok := alternate.(ast.BodyProgram); ok {
				for _, statement := range alternate.(ast.BodyProgram) {
					i.InterpretNode(statement, "")
				}
			}
		}
	}
}
