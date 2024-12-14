package interpreter

import "github.com/devnazir/ip-address-language/pkg/ast"

func (i *Interpreter) InterpretCallExpression(callExpression ast.CallExpression) (IntrerpretResult, ShouldReturn, error) {
	identName := (callExpression).Callee.(ast.Identifier).Name
	info := i.scopeResolver.ResolveScope(identName)
	arguments := (callExpression).Arguments
	res, shouldReturn, err := i.InterpretBodyFunction(info.Value.(ast.FunctionDeclaration), arguments)
	return res, shouldReturn, err
}
