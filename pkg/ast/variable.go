package ast

type VariableDeclaration struct {
	BaseNode
	Declarations   []VariableDeclarator
	Kind           string
	TypeAnnotation string
}

func (vd VariableDeclaration) GetLine() int {
	return vd.Declarations[0].Line
}

type VariableDeclarator struct {
	BaseNode
	Id   ASTNode
	Init ASTNode
}
