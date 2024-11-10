package parser

import lx "github.com/devnazir/gosh-script/internal/lexer"

type Parser struct {
	tokens []lx.Token
	lexer  lx.Lexer
	pos    int
}

var Precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
