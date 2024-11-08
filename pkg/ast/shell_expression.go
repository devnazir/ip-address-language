package ast

type ShellExpression struct {
	BaseNode
	Expression ASTNode
}

type EchoStatement struct {
	BaseNode
	Arguments []ASTNode
	Flags     []string
}

type LsStatement struct {
	BaseNode
}
