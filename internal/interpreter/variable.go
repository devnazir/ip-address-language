package interpreter

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/semantics"
	"github.com/devnazir/gosh-script/pkg/utils"
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

	if nodeVar.Declaration.Init == nil {
		typeAnnotation := nodeVar.TypeAnnotation
		return utils.InferDefaultValue(typeAnnotation)
	}

	declarationType := nodeVar.Declaration.Init.GetType()

	if declarationType == ast.SubShellTree {
		return i.InterpretSubShell(nodeVar.Declaration.Init.(ast.SubShell).Arguments)
	}

	if declarationType == ast.FunctionDeclarationTree {
		fnDeclaration := nodeVar.Declaration.Init.(ast.FunctionDeclaration)

		if !fnDeclaration.IsAnonymous {
			panic("Function declaration must be anonymous")
		}

		return fnDeclaration
	}

	if declarationType == ast.StringTemplateLiteralTree {
		stringTemplateLiteral := nodeVar.Declaration.Init.(ast.StringTemplateLiteral)
		var result string

		for _, part := range stringTemplateLiteral.Parts {
			result += i.InterpretBinaryExpr(part).(string)
		}

		return result
	}

	return i.InterpretBinaryExpr(nodeVar.Declaration.Init)
}
