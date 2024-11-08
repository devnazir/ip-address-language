package ast

type Literal struct {
	BaseNode
	Value interface{}
	Raw   string
}
