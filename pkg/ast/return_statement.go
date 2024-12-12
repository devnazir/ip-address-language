package ast

const ReturnStatementTree = "ReturnStatement"

type ReturnStatement struct {
	BaseNode
	Argument []ASTNode
}

func (r ReturnStatement) GetLine() int {
	return r.Line
}
