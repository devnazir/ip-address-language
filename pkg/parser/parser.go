package parser

import (
	"fmt"

	lx "github.com/devnazir/gosh-script/pkg/lexer"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func NewParser(tokens []lx.Token, lexer lx.Lexer) *Parser {
	return &Parser{
		tokens: tokens,
		lexer:  lexer,
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
		return p.tokens[len(p.tokens)-1]
	}

	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) Parse() ASTNode {
	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	return p.ParseProgram()
}

func (p *Parser) ParseProgram() Program {
	program := Program{
		BaseNode: BaseNode{
			Type:  "Program",
			Start: 0,
			End:   len(p.lexer.Source),
		},
		Body: []ASTNode{},
	}

	for p.pos < len(p.tokens) {
		switch p.peek().Type {
		case KEYWORD:
			if p.peek().Value == lx.VAR || p.peek().Value == lx.CONST {
				program.Body = append(program.Body, p.ParseVariableDeclaration())
			} else {
				oops.UnexpectedKeyword(p.peek())
			}
		case lx.SEMICOLON:
		case lx.COMMENT:
			p.next()
		case ILLEGAL:
			oops.IllegalToken(p.peek())
		case EOF:
			return program
		default:
			oops.UnexpectedToken(p.peek(), "")
		}
	}

	return program
}
