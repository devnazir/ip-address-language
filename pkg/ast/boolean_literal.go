package ast

const (
	BooleanLiteralTree = "BooleanLiteral"
)

type BooleanLiteral struct {
	BaseNode
	Value bool
	Raw   string
}
