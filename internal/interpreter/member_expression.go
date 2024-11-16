package interpreter

import (
	"fmt"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (i *Interpreter) InterpretMemberExpr(expr ast.MemberExpression) interface{} {
	object := i.EvaluateMemberExpr(expr.Object)
	property := i.EvaluateMemberExpr(expr.Property)
	computed := expr.Computed

	if computed {
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

	return nil
}

func (i *Interpreter) EvaluateMemberExpr(node ast.ASTNode) interface{} {
	switch node.(type) {
	case ast.Identifier:
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
