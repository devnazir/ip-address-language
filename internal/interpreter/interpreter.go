package interpreter

import (
	"fmt"

	"github.com/devnazir/gosh-script/pkg/ast"
)

var env = NewEnvironment()

func NewInterpreter() *Interpreter {
	return &Interpreter{Environment: env}
}

func (i *Interpreter) Interpret(p ast.ASTNode) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	program := p.(ast.Program)

	for _, nodeItem := range program.Body {
		switch nodeItem.(type) {
		case ast.VariableDeclaration:
			i.InterpretVariableDeclaration(nodeItem.(ast.VariableDeclaration))

		case ast.ShellExpression:
			i.InterpretShellExpression(InterpretShellExpression{
				expression:    nodeItem.(ast.ShellExpression),
				captureOutput: false,
			})
		}
	}
}
