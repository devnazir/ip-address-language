package ast

type FunctionDeclaration struct {
	BaseNode
	Identifier
	Body       []ASTNode
	Parameters []Identifier
}
