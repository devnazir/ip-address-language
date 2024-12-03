package ast

const (
	VariableDeclarationTree = "VariableDeclaration"
	VariableDeclaratorTree  = "VariableDeclarator"
)

type VariableDeclaration struct {
	BaseNode
	Declaration    VariableDeclarator
	Kind           string
	TypeAnnotation string
}

func (vd VariableDeclaration) GetLine() int {
	return vd.Declaration.Line
}

type VariableDeclarator struct {
	BaseNode
	Id   ASTNode
	Init ASTNode
}

func (vd VariableDeclarator) GetLine() int {
	return vd.Line
}
