package interpreter

import (
	"fmt"
	"runtime/debug"

	"github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/semantics"
)

func NewInterpreter() *Interpreter {
	symbolTable := semantics.NewSymbolTable()
	scopeResolver := semantics.NewScopeResolver(symbolTable)

	return &Interpreter{
		symbolTable:   symbolTable,
		scopeResolver: scopeResolver,
	}
}

func (i *Interpreter) Interpret(p ast.ASTNode) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
		}

		// utils.PrintJson(i.symbolTable.Scopes)
	}()

	program := p.(*ast.Program)
	entryPoint := program.EntryPoint

	for _, nodeItem := range program.Body {
		err := i.InterpretNode(nodeItem, entryPoint)
		if err != nil {
			panic(err)
		}
	}
}

func (i *Interpreter) InterpretNode(nodeItem ast.ASTNode, entryPoint string) error {
	switch nodeItem.GetType() {
	case ast.VariableDeclarationTree:
		i.InterpretVariableDeclaration((nodeItem).(ast.VariableDeclaration))

	case ast.ShellExpressionTree:
		i.InterpretShellExpression(InterpretShellExpression{
			expression:    (nodeItem).(ast.ShellExpression),
			captureOutput: false,
		})

	case ast.SubShellTree:
		res := i.InterpretSubShell((nodeItem).(ast.SubShell).Arguments)
		fmt.Printf("%v", res)

	case ast.AssignmentExpressionTree:
		i.InterpretAssigmentExpression((nodeItem).(ast.AssignmentExpression))

	case ast.SourceDeclarationTree:
		i.InterpretSourceDeclaration((nodeItem).(ast.SourceDeclaration), entryPoint)

	case ast.FunctionDeclarationTree:
		name := (nodeItem).(ast.FunctionDeclaration).Identifier.Name

		if name == "init" && len((nodeItem).(ast.FunctionDeclaration).Parameters) > 0 {
			return oops.RuntimeError(nodeItem, "init function cannot have parameters")
		}

		i.symbolTable.Insert((nodeItem).(ast.FunctionDeclaration).Identifier.Name, semantics.SymbolInfo{
			Kind:  lexer.KeywordFunc,
			Value: (nodeItem).(ast.FunctionDeclaration),
			Line:  (nodeItem).(ast.FunctionDeclaration).Line,
		})

	case ast.CallExpressionTree:
		info := i.scopeResolver.ResolveScope((nodeItem).(ast.CallExpression).Callee.(ast.Identifier).Name)
		arguments := (nodeItem).(ast.CallExpression).Arguments
		i.InterpretBodyFunction(info.Value.(ast.FunctionDeclaration), arguments)
	}

	// utils.PrintJson(i.symbolTable)

	return nil
}
