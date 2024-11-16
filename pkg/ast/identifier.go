package ast

type Identifier struct {
	BaseNode
	Name            string
	IsRestParameter bool
}

func (i Identifier) GetLine() int {
	return i.Line
}
