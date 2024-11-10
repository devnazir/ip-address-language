package ast

type NumberLiteral struct {
	BaseNode
	Value interface{}
	Raw   string
}
