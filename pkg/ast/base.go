package ast

const (
	BaseTree = "BaseNode"
)

type ASTNode interface {
	GetType() interface{}
	GetLine() int
}

type BodyProgram []ASTNode;

func (b BodyProgram) GetType() interface{} {
	return b
}

func (b BodyProgram) GetLine() int {
	return 0
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
