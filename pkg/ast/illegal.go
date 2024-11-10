package ast

type Illegal struct {
	BaseNode
	Value string
	Raw   string
}
