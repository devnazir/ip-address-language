package ast

const (
	ProgramTree = "Program"
)

type Program struct {
	BaseNode
	Body       []ASTNode
	EntryPoint string
}
