package ast

type StringLiteral struct {
	BaseNode
	Value string
	Raw   string
}
