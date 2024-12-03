package ast

const (
	IllegalTree = "Illegal"
)

type Illegal struct {
	BaseNode
	Value string
	Raw   string
}
