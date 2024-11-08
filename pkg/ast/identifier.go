package ast

type Identifier struct {
	BaseNode
	Name string
}

func (i Identifier) GetLine() int {
	return i.Line
}
