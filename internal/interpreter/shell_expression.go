package interpreter

import "github.com/devnazir/gosh-script/pkg/ast"

func (i *Interpreter) InterpretShellExpression(params InterpretShellExpression) interface{} {
	nodeShell := params.expression
	captureOutput := params.captureOutput
	expression := nodeShell.Expression

	var result interface{}

	switch expression.(type) {
	case ast.EchoStatement:
		res := i.IntrepretEchoStmt(IntrepretEchoStmt{
			expression:    expression.(ast.EchoStatement),
			captureOutput: captureOutput,
		})
		result = res
	}

	if captureOutput {
		return result
	}

	return nil
}
