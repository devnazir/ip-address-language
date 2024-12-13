package interpreter

import (
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (i *Interpreter) InterpretReturnStatement(node ast.ReturnStatement) (interface{}, error) {
	var result []interface{}

	for _, arg := range node.Arguments {
		interpretedArg, _, err := i.InterpretNode(arg, "")

		if err != nil {
			return nil, err
		}

		result = append(result, interpretedArg)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return result[0], nil
}
