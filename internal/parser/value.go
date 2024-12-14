package parser

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
)

type Parser struct {
	tokens []lx.Token
	pos    int
}

var Precedence = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"%": 2,
}

var ComparisonPrecedence = map[string]int{
	"||": 1,
	"&&": 2,
	"==": 3, "!=": 3,
	">": 4, "<": 4, ">=": 4, "<=": 4,

	"+": 5,
	"-": 5,
	"*": 6,
	"/": 6,
	"%": 6,
}
