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

func (e EchoStatement) GetLine() int {
	return e.Line
}

type SubShell struct {
	BaseNode
	Arguments ASTNode
}
