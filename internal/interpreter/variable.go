package interpreter

import (
	"fmt"

	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
	"github.com/devnazir/ip-address-language/pkg/semantics"
	"github.com/devnazir/ip-address-language/pkg/utils"
)

func (i *Interpreter) InterpretVariableDeclaration(nodeVar ast.VariableDeclaration) error {
	name := nodeVar.Declaration.Id.(ast.Identifier).Name

	if i.symbolTable.Exists(name) {
		return oops.SyntaxError(nodeVar, fmt.Sprintf("Variable %s already declared", name))
	}

	value, _ := i.EvaluateVariableInit(nodeVar)

	symbolInfo := &semantics.SymbolInfo{
		Kind:           nodeVar.Kind,
		Value:          value,
		Line:           nodeVar.Line,
		TypeAnnotation: nodeVar.TypeAnnotation,
	}

	i.symbolTable.Insert(name, *symbolInfo)
	return nil
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
			expr := i.InterpretBinaryExpr(part, true)
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

			result = append(result, i.InterpretBinaryExpr(element, true))
		}

		return result, ast.ArrayExpressionTree
	}

	if declarationType == ast.ObjectExpressionTree {
		objectExpression := nodeVar.Declaration.Init.(ast.ObjectExpression)
		result := make(map[string]interface{})

		for _, property := range objectExpression.Properties {
			key := property.Key
			value := i.InterpretBinaryExpr(property.Value, true)

			result[key] = value
		}

		return result, ast.ObjectExpressionTree
	}

	if declarationType == ast.CallExpressionTree {
		res, _, _ := i.InterpretNode(nodeVar.Declaration.Init, "")
		return res, ast.CallExpressionTree
	}

	return i.InterpretBinaryExpr(nodeVar.Declaration.Init, true), ast.BinaryExpressionTree
}
