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

type Parser struct {
	tokens []lx.Token
	lexer  lx.Lexer
	pos    int
}

var TokenMap = lx.TokenMap()

var Precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
