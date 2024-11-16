package interpreter

import (
	"fmt"
	"runtime/debug"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/semantics"
)

func (i *Interpreter) InterpretBodyFunction(p ast.FunctionDeclaration, args []ast.ASTNode) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
		}
	}()

	i.scopeResolver.EnterScope()
	idxOfRestParam := len(p.Parameters) - 1
	lenParams := len(p.Parameters)
	hasRestParam := false

	for idx, param := range p.Parameters {
		isRest := param.IsRestParameter
		if isRest {
			hasRestParam = true
			idxOfRestParam = idx
			lenParams -= 1
		}
	}

	if !hasRestParam {
		idxOfRestParam = len(p.Parameters)
	}

	if len(args[:idxOfRestParam]) < lenParams {
		panic("Function not called with enough arguments")
	}

	if len(args[:idxOfRestParam]) > lenParams {
		panic("Function called with too many arguments")
	}

	restArgs := args[idxOfRestParam:]
	restArgsValues := make([]interface{}, len(restArgs))

	for idx, arg := range restArgs {
		switch arg := arg.(type) {
		case ast.StringLiteral:
			restArgsValues[idx] = arg.Value
		case ast.NumberLiteral:
			restArgsValues[idx] = arg.Value
		case ast.Identifier:
			name := arg.Name
			info := i.scopeResolver.ResolveScope(name)

			restArgsValues[idx] = info.Value
		}
	}

	for idx, param := range p.Parameters {
		name := param.Name
		isRest := param.IsRestParameter

		var value interface{}

		if isRest && idx != len(p.Parameters)-1 {
			panic("Rest parameter must be the last parameter")
		}

		if isRest {
			i.symbolTable.Insert(name, semantics.SymbolInfo{
				Value: restArgsValues,
			})
			break
		}

		switch arg := args[idx].(type) {
		case ast.StringLiteral:
			value = arg.Value
		case ast.NumberLiteral:
			value = arg.Value
		case ast.Identifier:
			name := arg.Name
			info := i.scopeResolver.ResolveScope(name)

			value = info.Value
		}

		i.symbolTable.Insert(name, semantics.SymbolInfo{
			Value: value,
		})
	}

	for _, nodeItem := range p.Body {
		i.InterpretNode(nodeItem, "")
	}

	i.scopeResolver.ExitScope()
}
