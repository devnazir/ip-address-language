package parser

import (
	"fmt"
	"runtime/debug"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func NewParser(tokens *[]lx.Token) *Parser {
	return &Parser{
		tokens: *tokens,
		pos:    0,
	}
}

func (p *Parser) peek() lx.Token {
	if p.pos >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[p.pos]
}

func (p *Parser) next() lx.Token {
	if p.pos >= len(p.tokens) {
		return lx.Token{Type: lx.TokenEOF}
	}

	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) Parse() *ast.Program {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
		}
	}()

	program := p.ParseProgram()

	return program
}

func (p *Parser) TokenTypeIs(t lx.TokenType) bool {
	return p.peek().Type == t
}

func (p *Parser) TokenValueIs(v string) bool {
	return p.peek().Value == v
}
