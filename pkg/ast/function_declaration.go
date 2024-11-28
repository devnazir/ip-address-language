package ast

type FunctionDeclaration struct {
	BaseNode
	Identifier
	Body       []ASTNode
	Parameters []Identifier
	IsAnonymous bool
}

func (f FunctionDeclaration) GetLine() int {
	return f.Line
}
