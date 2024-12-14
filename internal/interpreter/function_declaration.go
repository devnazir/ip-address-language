package interpreter

import (
	"fmt"

	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
	"github.com/devnazir/ip-address-language/pkg/semantics"
)

func (i *Interpreter) processFnArgument(arg ast.ASTNode) string {
	switch arg.GetType() {
	case ast.StringLiteralTree:
		return arg.(ast.StringLiteral).Value
	case ast.NumberLiteralTree:
		return fmt.Sprintf("%v", arg.(ast.NumberLiteral).Value)
	case ast.IdentifierTree:
		name := arg.(ast.Identifier).Name
		info := i.scopeResolver.ResolveScope(name)

		return fmt.Sprintf("%v", info.Value)
	}

	return ""
}

func (i *Interpreter) InterpretBodyFunction(p ast.FunctionDeclaration, args []ast.ASTNode) (IntrerpretResult, ShouldReturn, error) {
	i.scopeResolver.EnterScope()

	idxOfRestParams, err := i.getRestParamIndex(&p)

	if err != nil {
		panic(err)
	}

	validateParameter(&p, args, idxOfRestParams)
	restArguments := i.getRestArguments(args, idxOfRestParams)

	for idx, param := range p.Parameters {
		name := param.Name
		isRest := param.IsRestParameter

		var value = i.processFnArgument(args[idx])

		if isRest {
			i.symbolTable.Insert(name, semantics.SymbolInfo{
				Value: restArguments,
			})
			break
		}

		i.symbolTable.Insert(name, semantics.SymbolInfo{
			Value: value,
		})
	}

	for _, nodeItem := range p.Body {
		res, shouldReturn, err := i.InterpretNode(nodeItem, "")
		if err != nil {
			return nil, false, err
		}

		if shouldReturn {
			return res, shouldReturn, nil
		}
	}

	i.scopeResolver.ExitScope()

	return nil, false, nil
}

func (i *Interpreter) getRestParamIndex(functionDecl *ast.FunctionDeclaration) (int, error) {
	var item ast.Identifier
	index := -1

	if len(functionDecl.Parameters) == 0 {
		return index, nil
	}

	if len(functionDecl.Parameters) > 1 {
		index = len(functionDecl.Parameters) - 1
		item = functionDecl.Parameters[index]
	} else {
		item = functionDecl.Parameters[0]
		index = 0
	}

	validateMiddleRestParameters := func() error {
		if len(functionDecl.Parameters) > 1 {
			for _, param := range functionDecl.Parameters[1:index] {
				if param.IsRestParameter {
					return fmt.Errorf("Rest parameter must be last")
				}
			}
		}

		return nil
	}

	if !item.IsRestParameter {

		if index > 0 {
			// check the first item
			firstItem := functionDecl.Parameters[0]
			if firstItem.IsRestParameter {
				return index, fmt.Errorf("Rest parameter must be last")
			}
		}

		err := validateMiddleRestParameters()
		if err != nil {
			return index, err
		}

		index = -1
	}

	// check if there is duplicate rest parameter
	err := validateMiddleRestParameters()
	if err != nil {
		return index, fmt.Errorf("Duplicate rest parameter")
	}

	if item.IsRestParameter {
		return index, nil
	}

	return index, nil
}

func validateParameter(functionDecl *ast.FunctionDeclaration, args []ast.ASTNode, restParamIndex int) {
	hasRestParam := restParamIndex != -1
	paramCount := len(functionDecl.Parameters)

	argsLimit := len(args)
	if hasRestParam {
		argsLimit = restParamIndex
		paramCount--
	}

	argsCount := len(args[:argsLimit])

	if argsCount < paramCount {
		oops.FunctionNotCalledWithEnoughArgumentsError(*functionDecl, paramCount, argsCount)
	}
	if argsCount > paramCount {
		oops.FunctionCalledWithTooManyArgumentsError(*functionDecl, paramCount, argsCount)
	}
}

func (i *Interpreter) getRestArguments(args []ast.ASTNode, restParamIndex int) []interface{} {
	if len(args) == 0 || restParamIndex == -1 {
		return []interface{}{}
	}

	restArgs := args[restParamIndex:]
	restArgsValues := make([]interface{}, len(restArgs))

	for idx, arg := range restArgs {
		switch arg.GetType() {
		case ast.StringLiteralTree:
			restArgsValues[idx] = arg.(ast.StringLiteral).Value
		case ast.NumberLiteralTree:
			restArgsValues[idx] = arg.(ast.NumberLiteral).Value
		case ast.IdentifierTree:
			name := arg.(ast.Identifier).Name
			info := i.scopeResolver.ResolveScope(name)

			restArgsValues[idx] = info.Value
		}
	}

	return restArgsValues
}
