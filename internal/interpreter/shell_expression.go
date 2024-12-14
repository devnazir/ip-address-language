package interpreter

import (
	"fmt"

	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/semantics"
)

func (i *Interpreter) InterpretShellExpression(params InterpretShellExpression) interface{} {
	nodeShell := params.expression
	captureOutput := params.captureOutput
	expression := nodeShell.Expression

	var result interface{}

	switch expression.GetType() {
	case ast.EchoStatementTree:
		argsString := ""
		flags := ""
		for _, arg := range expression.(ast.EchoStatement).Arguments {
			argsString += i.processArgument(arg) + " "
		}

		for _, flag := range expression.(ast.EchoStatement).Flags {
			flags += flag + " "
		}

		address := i.symbolTable.MakeAddress(semantics.SymbolInfo{
			Value: map[string]string{
				"args":  argsString,
				"flags": flags,
			},
		})

		if _, ok := i.symbolTable.Address[address]; ok {
			result = i.symbolTable.Address[address].Value
		} else {
			result = i.IntrepretEchoStmt(IntrepretEchoStmt{
				expression: expression.(ast.EchoStatement),
			})

			i.symbolTable.InsertAddress(address, semantics.SymbolInfo{
				Value: result,
			})
		}

		if captureOutput {
			return result
		}

		fmt.Printf("%v", result)
	}

	return nil
}
