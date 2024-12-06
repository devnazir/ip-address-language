package ast

const ArrayExpressionTree = "ArrayExpression"

type ArrayExpression struct {
	BaseNode
	Elements []ASTNode
}
