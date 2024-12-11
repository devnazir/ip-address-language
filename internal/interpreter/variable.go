package interpreter

import (
	"fmt"

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

	value, _ := i.EvaluateVariableInit(nodeVar)

	symbolInfo := &semantics.SymbolInfo{
		Kind:           nodeVar.Kind,
		Value:          value,
		Line:           nodeVar.Line,
		TypeAnnotation: nodeVar.TypeAnnotation,
	}

	i.symbolTable.Insert(name, *symbolInfo)
}

func (i *Interpreter) EvaluateVariableInit(nodeVar ast.VariableDeclaration) (interface{}, string) {

	if nodeVar.Declaration.Init == nil {
		typeAnnotation := nodeVar.TypeAnnotation
		return utils.InferDefaultValue(typeAnnotation), ""
	}

	declarationType := nodeVar.Declaration.Init.GetType()

	if declarationType == ast.SubShellTree {
		return i.InterpretSubShell(nodeVar.Declaration.Init.(ast.SubShell).Arguments), ast.SubShellTree
	}

	if declarationType == ast.FunctionDeclarationTree {
		fnDeclaration := nodeVar.Declaration.Init.(ast.FunctionDeclaration)

		if !fnDeclaration.IsAnonymous {
			panic(oops.SyntaxError(fnDeclaration, "Named function declaration not allowed"))
		}

		return fnDeclaration, ast.FunctionDeclarationTree
	}

	if declarationType == ast.StringTemplateLiteralTree {
		stringTemplateLiteral := nodeVar.Declaration.Init.(ast.StringTemplateLiteral)
		var result string

		for _, part := range stringTemplateLiteral.Parts {
			expr := i.InterpretBinaryExpr(part)
			result += fmt.Sprintf("%v", expr)
		}

		return result, ast.StringTemplateLiteralTree
	}

	if declarationType == ast.ArrayExpressionTree {
		arrayExpression := nodeVar.Declaration.Init.(ast.ArrayExpression)
		var result []interface{}

		for _, element := range arrayExpression.Elements {

			if element.GetType() == ast.ArrayExpressionTree {
				node, _ := i.EvaluateVariableInit(ast.VariableDeclaration{
					Declaration: ast.VariableDeclarator{
						Init: element,
					},
				})
				result = append(result, node)
				continue
			}

			result = append(result, i.InterpretBinaryExpr(element))
		}

		return result, ast.ArrayExpressionTree
	}

	if declarationType == ast.ObjectExpressionTree {
		objectExpression := nodeVar.Declaration.Init.(ast.ObjectExpression)
		result := make(map[string]interface{})

		for _, property := range objectExpression.Properties {
			key := property.Key
			value := i.InterpretBinaryExpr(property.Value)

			result[key] = value
		}

		return result, ast.ObjectExpressionTree
	}

	return i.InterpretBinaryExpr(nodeVar.Declaration.Init), ast.BinaryExpressionTree
}
