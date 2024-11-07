package node

type BaseNode struct {
	Type  string
	Start int
	End   int
	Line  int
}

type ASTNode interface{}

type Program struct {
	BaseNode
	Body []ASTNode
}

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

type Identifier struct {
	BaseNode
	Name string
}

func (i Identifier) GetLine() int {
	return i.Line
}

type Literal struct {
	BaseNode
	Value interface{}
	Raw   string
}

type BinaryExpression struct {
	BaseNode
	Left     ASTNode
	Operator string
	Right    ASTNode
}

type ShellExpression struct {
	BaseNode
	Expression ASTNode
}

type EchoStatement struct {
	BaseNode
	Arguments []ASTNode
	Flags     []string
}

type AssignmentExpression struct {
	Identifier
	Expression ASTNode
}

type SourceDeclaration struct {
	BaseNode
	Sources []ASTNode
}
