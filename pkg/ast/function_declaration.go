package ast

const (
	FunctionDeclarationTree = "FunctionDeclaration"
)

type FunctionDeclaration struct {
	BaseNode
	Identifier  Identifier
	Body        []ASTNode
	Parameters  []Identifier
	IsAnonymous bool
}

func (f FunctionDeclaration) GetLine() int {
	return f.Line
}
