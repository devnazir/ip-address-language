package ast

type BinaryExpression struct {
	BaseNode
	Left     ASTNode
	Operator string
	Right    ASTNode
}
