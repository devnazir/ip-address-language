package interpreter

import (
	"reflect"

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

	if _, ok := nodeVar.Declaration.Init.(ast.FunctionDeclaration); ok {
		fnDeclaration := nodeVar.Declaration.Init.(ast.FunctionDeclaration)

		if !fnDeclaration.IsAnonymous {
			panic("Function declaration must be anonymous")
		}

		return fnDeclaration
	}

	if _, ok := nodeVar.Declaration.Init.(ast.StringTemplateLiteral); ok {
		stringTemplateLiteral := nodeVar.Declaration.Init.(ast.StringTemplateLiteral)
		var result string

		for _, part := range stringTemplateLiteral.Parts {

			if reflect.TypeOf(part) == reflect.TypeOf(ast.Identifier{}) {
				// add space between variables
				result += " "
			}

			result += i.InterpretBinaryExpr(part).(string)
		}

		return result
	}

	return i.InterpretBinaryExpr(nodeVar.Declaration.Init)
}
