package ast

const (
	ShellExpressionTree = "ShellExpression"
	EchoStatementTree   = "EchoStatement"
	SubShellTree        = "SubShell"
)

type ShellExpression struct {
	BaseNode
	Expression ASTNode
}

func (s ShellExpression) GetLine() int {
	return s.Line
}

func (s ShellExpression) GetType() interface{} {
	return s.Type
}

type EchoStatement struct {
	BaseNode
	Arguments []ASTNode
	Flags     []string
}

func (e EchoStatement) GetLine() int {
	return e.Line
}

func (e EchoStatement) GetType() interface{} {
	return e.Type
}

type SubShell struct {
	BaseNode
	Arguments string
}

func (s SubShell) GetLine() int {
	return s.Line
}

func (s SubShell) GetType() interface{} {
	return s.Type
}
