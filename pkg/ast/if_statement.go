package ast

const IFStatementTree = "IfStatement"

type Condition ASTNode

type IfStatement struct {
	BaseNode
	Condition  Condition
	Consequent interface{}
	Alternate  interface{}
}

func (i IfStatement) GetLine() int {
	return i.Line
}
