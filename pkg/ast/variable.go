package ast

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
