package ast

const ObjectExpressionTree = "ObjectExpression"

type Property struct {
	BaseNode
	Key   string
	Value ASTNode
}

type ObjectExpression struct {
	BaseNode
	Properties []Property
}
