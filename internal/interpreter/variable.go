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

	value := i.EvaluateVariableInit(nodeVar)

	i.symbolTable.Insert(name, semantics.SymbolInfo{
		Kind:           nodeVar.Kind,
		Value:          value,
		Line:           nodeVar.Line,
		TypeAnnotation: nodeVar.TypeAnnotation,
	})
}

func (i *Interpreter) EvaluateVariableInit(nodeVar ast.VariableDeclaration) interface{} {
	if _, ok := nodeVar.Declaration.Init.(ast.SubShell); ok {
		return i.InterpretSubShell(nodeVar.Declaration.Init.(ast.SubShell).Arguments.(string))
	}

	return i.InterpretBinaryExpr(nodeVar.Declaration.Init)
}
