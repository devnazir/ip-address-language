package ast

const (
	BaseTree = "BaseNode"
)

type ASTNode interface {
	GetType() interface{}
	GetLine() int
}

type BaseNode struct {
	Type    string
	Start   int
	End     int
	Line    int
}

func (b BaseNode) GetType() interface{} {
	return b.Type
}

func (b BaseNode) GetLine() int {
	return b.Line
}
