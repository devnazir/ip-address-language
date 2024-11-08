package ast

type AssignmentExpression struct {
	Identifier
	Expression ASTNode
}
