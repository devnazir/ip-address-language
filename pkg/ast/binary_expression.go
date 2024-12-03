package ast

const (
	BinaryExpressionTree = "BinaryExpression"
)

type BinaryExpression struct {
	BaseNode
	Left     ASTNode
	Operator string
	Right    ASTNode
}

func (b BinaryExpression) GetLine() int {
	return b.Line
}

func (b BinaryExpression) GetType() interface{} {
	return b.Type
}
