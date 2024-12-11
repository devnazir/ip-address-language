package interpreter

import (
	"fmt"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (i *Interpreter) InterpretMemberExpr(expr ast.MemberExpression) interface{} {
	computed := expr.Computed

	if computed {
		object := i.EvaluateMemberExpr(expr.Object, computed)
		property := i.EvaluateMemberExpr(expr.Property, computed)
		indexable, ok := object.([]interface{})
		if !ok {
			panic(oops.RuntimeError(expr.Property, "Not indexable"))
		}

		index, ok := property.(int)
		if !ok {
			panic(oops.SyntaxError(expr.Property, "Invalid index"))
		}

		if len(indexable) <= index {
			panic(oops.RuntimeError(expr.Property, fmt.Sprintf("Index %v out of range", index)))
		}

		value := indexable[index]
		return value
	}

	object := i.EvaluateMemberExpr(expr.Object, true)
	indexable, ok := object.(map[string]interface{})
	if !ok {
		panic(oops.RuntimeError(expr.Property, "Not indexable"))
	}

	property := i.EvaluateMemberExpr(expr.Property, computed)
	value, ok := indexable[property.(string)]
	if !ok {
		panic(oops.RuntimeError(expr.Property, fmt.Sprintf("Property %v not found", property)))
	}

	return value
}

func (i *Interpreter) EvaluateMemberExpr(node ast.ASTNode, computed bool) interface{} {
	switch node.GetType() {
	case ast.IdentifierTree:
		if !computed {
			return node.(ast.Identifier).Name
		}

		name := node.(ast.Identifier).Name
		info := i.scopeResolver.ResolveScope(name)

		return info.Value

	case ast.NumberLiteralTree:
		return node.(ast.NumberLiteral).Value

	case ast.StringLiteralTree:
		return node.(ast.StringLiteral).Value

	case ast.CallExpressionTree:
		info := i.scopeResolver.ResolveScope((node).(ast.CallExpression).Callee.(ast.Identifier).Name)
		return info.Value

	case ast.MemberExpressionTree:
		return i.InterpretMemberExpr(node.(ast.MemberExpression))

	default:
		panic(oops.RuntimeError(node, "Invalid member expression"))
	}
}
