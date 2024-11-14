package interpreter

import (
	"fmt"
	"runtime/debug"

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
			debug.PrintStack()
		}
	}()

	program := p.(*ast.Program)
	entryPoint := program.EntryPoint

	for _, nodeItem := range program.Body {
		InterpretNode(i, nodeItem, entryPoint)
	}
}

func (i *Interpreter) InterpretFunctionDeclaration(p ast.FunctionDeclaration) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
		}
	}()

	for _, nodeItem := range p.Body {
		InterpretNode(i, nodeItem, "")
	}
}

func InterpretNode(i *Interpreter, nodeItem ast.ASTNode, entryPoint string) {

	switch (nodeItem).(type) {
	case ast.VariableDeclaration:
		i.InterpretVariableDeclaration((nodeItem).(ast.VariableDeclaration))

	case ast.ShellExpression:
		i.InterpretShellExpression(InterpretShellExpression{
			expression:    (nodeItem).(ast.ShellExpression),
			captureOutput: false,
		})

	case ast.SubShell:
		res := i.InterpretSubShell((nodeItem).(ast.SubShell).Arguments.(string))
		fmt.Printf("%v", res)

	case ast.AssignmentExpression:
		i.InterpretAssigmentExpression((nodeItem).(ast.AssignmentExpression))

	case ast.SourceDeclaration:
		i.InterpretSourceDeclaration((nodeItem).(ast.SourceDeclaration).Sources, entryPoint)
	}
}
