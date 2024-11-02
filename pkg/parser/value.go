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
	LEFT_PAREN     = lx.LPAREN
	RIGHT_PAREN    = lx.RPAREN
	OPERATOR       = lx.OPERATOR
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
	Declarations   []VariableDeclarator
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

type BinaryExpression struct {
	BaseNode
	Left     ASTNode
	Operator string
	Right    ASTNode
}

type AssignmentExpression struct {
	Identifier
	Expression ASTNode
}

var TokenMap = lx.TokenMap()

var Precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
