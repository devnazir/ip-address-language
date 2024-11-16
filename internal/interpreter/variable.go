package interpreter

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/semantics"
)

func (i *Interpreter) InterpretVariableDeclaration(nodeVar ast.VariableDeclaration) {
	name := nodeVar.Declaration.Id.(ast.Identifier).Name

	if i.symbolTable.Exists(name) {
		oops.DuplicateIdentifierError(nodeVar)
	}

	if _, ok := nodeVar.Declaration.Init.(ast.SubShell); ok {
		res := i.InterpretSubShell(nodeVar.Declaration.Init.(ast.SubShell).Arguments.(string))
		i.symbolTable.Insert(name, semantics.SymbolInfo{
			Type:  "",
			Value: res,
		})
		return
	}

	value := i.InterpretBinaryExpr(nodeVar.Declaration.Init)
	i.symbolTable.Insert(name, semantics.SymbolInfo{
		Type:  "",
		Value: value,
	})
}
