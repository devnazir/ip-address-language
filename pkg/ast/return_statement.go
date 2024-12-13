package ast

const ReturnStatementTree = "ReturnStatement"

type ReturnStatement struct {
	BaseNode
	Arguments []ASTNode
}

func (r ReturnStatement) GetLine() int {
	return r.Line
}
