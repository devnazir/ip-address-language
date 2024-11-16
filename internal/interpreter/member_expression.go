package interpreter

import (
	"fmt"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (i *Interpreter) InterpretMemberExpr(expr ast.MemberExpression) interface{} {
	computed := expr.Computed
	object := i.EvaluateMemberExpr(expr.Object, true)

	if computed {
		property := i.EvaluateMemberExpr(expr.Property, computed)
		indexable, ok := object.([]interface{})
		if !ok {
			panic("Not indexable")
		}

		index, ok := property.(int)
		if !ok {
			panic(fmt.Sprintf("Invalid index: %v", property))
		}

		if len(indexable) <= index {
			panic(fmt.Sprintf("Index %v out of range", index))
		}

		value := indexable[index]
		return value
	}

	indexable, ok := object.(map[string]interface{})
	if !ok {
		panic("Not indexable")
	}

	property := i.EvaluateMemberExpr(expr.Property, computed)
	value, ok := indexable[property.(string)]
	if !ok {
		panic(fmt.Sprintf("Property %v not found", expr.Property.(ast.Identifier).Name))
	}

	return value
}

func (i *Interpreter) EvaluateMemberExpr(node ast.ASTNode, computed bool) interface{} {
	switch node.(type) {
	case ast.Identifier:
		if !computed {
			return node.(ast.Identifier).Name
		}

		name := node.(ast.Identifier).Name
		info := i.scopeResolver.ResolveScope(name)

		return info.Value

	case ast.NumberLiteral:
		return node.(ast.NumberLiteral).Value

	case ast.StringLiteral:
		return node.(ast.StringLiteral).Value

	case ast.MemberExpression:
		return i.InterpretMemberExpr(node.(ast.MemberExpression))

	default:
		panic("Invalid member expression")
	}
}
