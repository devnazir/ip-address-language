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
	}()

	program := p.(*ast.Program)
	entryPoint := program.EntryPoint

	for _, nodeItem := range program.Body {
		_, _, err := i.InterpretNode(nodeItem, entryPoint)
		if err != nil {
			panic(err)
		}
	}
}

type ShouldReturn bool
type IntrerpretResult interface{}

func (i *Interpreter) InterpretNode(nodeItem ast.ASTNode, entryPoint string) (IntrerpretResult, ShouldReturn, error) {

	switch nodeItem.GetType() {
	case ast.VariableDeclarationTree:
		err := i.InterpretVariableDeclaration((nodeItem).(ast.VariableDeclaration))
		return nil, false, err

	case ast.ShellExpressionTree:
		res := i.InterpretShellExpression(InterpretShellExpression{
			expression:    (nodeItem).(ast.ShellExpression),
			captureOutput: false,
		})
		return res, false, nil

	case ast.SubShellTree:
		res := i.InterpretSubShell((nodeItem).(ast.SubShell).Arguments)
		fmt.Printf("%v", res)
		return res, false, nil

	case ast.AssignmentExpressionTree:
		err := i.InterpretAssigmentExpression((nodeItem).(ast.AssignmentExpression))
		return nil, false, err

	case ast.SourceDeclarationTree:
		err := i.InterpretSourceDeclaration((nodeItem).(ast.SourceDeclaration), entryPoint)
		return nil, false, err

	case ast.FunctionDeclarationTree:
		name := (nodeItem).(ast.FunctionDeclaration).Identifier.Name

		if name == "init" && len((nodeItem).(ast.FunctionDeclaration).Parameters) > 0 {
			return nil, false, oops.RuntimeError(nodeItem, "init function cannot have parameters")
		}

		i.symbolTable.Insert((nodeItem).(ast.FunctionDeclaration).Identifier.Name, semantics.SymbolInfo{
			Kind:  lexer.KeywordFunc,
			Value: (nodeItem).(ast.FunctionDeclaration),
			Line:  (nodeItem).(ast.FunctionDeclaration).Line,
		})

		return nil, false, nil

	case ast.CallExpressionTree:
		res, shouldReturn, err := i.InterpretCallExpression((nodeItem).(ast.CallExpression))
		return res, shouldReturn, err

	case ast.ReturnStatementTree:
		res, err := i.InterpretReturnStatement((nodeItem).(ast.ReturnStatement))
		return res, true, err

	case ast.IFStatementTree:
		i.scopeResolver.EnterScope()
		res, shouldReturn, err := i.InterpretIfStatement((nodeItem).(ast.IfStatement))
		i.scopeResolver.ExitScope()
		return res, shouldReturn, err

	case ast.MemberExpressionTree:
		i.InterpretMemberExpr(nodeItem.(ast.MemberExpression))
		return nil, false, nil

	case ast.StringLiteralTree:
		return (nodeItem).(ast.StringLiteral).Value, false, nil

	case ast.NumberLiteralTree:
		return (nodeItem).(ast.NumberLiteral).Value, false, nil

	case ast.IdentifierTree:
		return i.resolveIdentifier((nodeItem).(ast.Identifier)), false, nil

	}

	// utils.PrintJson(i.symbolTable)

	return nil, false, oops.RuntimeError(nodeItem, "Unknown node type")
}

func (i *Interpreter) resolveIdentifier(identifier ast.Identifier) interface{} {
	info := i.scopeResolver.ResolveScope(identifier.Name)
	return info.Value
}
