package ast

type Program struct {
	BaseNode
	Body       []ASTNode
	EntryPoint string
}
