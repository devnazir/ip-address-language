package ast

type Int64Literal struct {
	BaseNode
	Value int64
	Raw   string
}
