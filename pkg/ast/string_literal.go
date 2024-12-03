package ast

const (
	StringLiteralTree         = "StringLiteral"
	StringTemplateLiteralTree = "StringTemplateLiteral"
)

type StringLiteral struct {
	BaseNode
	Value string
	Raw   string
}

func (s StringLiteral) GetLine() int {
	return s.Line
}

type StringTemplateLiteral struct {
	BaseNode
	Parts []ASTNode
}

func (s StringTemplateLiteral) GetLine() int {
	return s.Line
}
