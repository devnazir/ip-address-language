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

func (i *Interpreter) InterpretIfStatement(node ast.IfStatement) (IntrerpretResult, ShouldReturn, error) {
	condition := i.InterpretBinaryExpr(node.Condition, false)
	consequent := node.Consequent

	if condition == nil {
		return nil, false, oops.RuntimeError(node.Condition, "Condition must be a boolean value")
	}

	if !isBooleanValue(condition) {
		return nil, false, oops.RuntimeError(node.Condition, "Condition must be a boolean value")
	}

	if condition.(bool) {
		if _, ok := consequent.(ast.BodyProgram); ok {
			for _, statement := range consequent.(ast.BodyProgram) {
				res, shouldReturn, err := i.InterpretNode(statement, "")
				if err != nil {
					return nil, false, err
				}

				if shouldReturn {
					return res, true, nil
				}
			}
		}
	} else {
		alternate := node.Alternate

		if reflect.TypeOf(alternate) == reflect.TypeOf(ast.IfStatement{}) {
			return i.InterpretIfStatement(alternate.(ast.IfStatement))
		} else {
			if _, ok := alternate.(ast.BodyProgram); ok {
				for _, statement := range alternate.(ast.BodyProgram) {
					res, shouldReturn, err := i.InterpretNode(statement, "")
					if err != nil {
						return nil, false, err
					}

					if shouldReturn {
						return res, true, nil
					}
				}
			}
		}
	}

	return nil, false, nil
}
