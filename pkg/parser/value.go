package parser

import lx "github.com/devnazir/gosh-script/pkg/lexer"

const (
	KEYWORD        = lx.KEYWORD
	IDENTIFIER     = lx.IDENTIFIER
	NUMBER         = lx.NUMBER
	STRING         = lx.STRING
	ILLEGAL        = lx.ILLEGAL
	EOF            = lx.EOF
	PRIMITIVE_TYPE = lx.PRIMITIVE_TYPE
)

type BaseNode struct {
	Type  string
	Start int
	End   int
}

type ASTNode interface{}

type Parser struct {
	tokens []lx.Token
	lexer  lx.Lexer
	pos    int
}

type Program struct {
	BaseNode
	Body []ASTNode
}

type VariableDeclaration struct {
	BaseNode
	Declarations   []ASTNode
	Kind           string
	TypeAnnotation string
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

type Literal struct {
	BaseNode
	Value interface{}
	Raw   string
}

var TokenMap = lx.TokenMap()
