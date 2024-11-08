package ast

type ASTNode interface{}

type BaseNode struct {
	Type  string
	Start int
	End   int
	Line  int
}
