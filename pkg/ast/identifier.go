package ast

const (
	IdentifierTree = "Identifier"
)

type Identifier struct {
	BaseNode
	Name            string
	Raw             string
	IsRestParameter bool
}

func (i Identifier) GetLine() int {
	return i.Line
}

func (i Identifier) GetStart() int {
	return i.Start
}

func (i Identifier) GetEnd() int {
	return i.End
}
