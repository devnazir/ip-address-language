package ast

type StringLiteral struct {
	BaseNode
	Value string
	Raw   string
}

type StringTemplateLiteral struct {
	BaseNode
	Parts []ASTNode
}
