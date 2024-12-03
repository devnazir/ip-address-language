package ast

const (
	NumberLiteralTree = "NumberLiteral"
)

type NumberLiteral struct {
	BaseNode
	Value interface{}
	Raw   string
}

func (n NumberLiteral) GetLine() int {
	return n.Line
}
