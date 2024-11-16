package ast

type AssignmentExpression struct {
	Identifier
	Expression ASTNode
}

type CallExpression struct {
	BaseNode
	Callee    ASTNode
	Arguments []ASTNode
}
