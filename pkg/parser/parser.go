package parser

import (
	"fmt"
	"strconv"

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
			if p.peek().Value == "var" {
				program.Body = append(program.Body, p.ParseVariableDeclaration())
			} else {
				oops.UnexpectedKeyword(p.peek())
			}
		case ILLEGAL:
			oops.IllegalToken(p.peek())
		case EOF:
			return program
		default:
			oops.UnexpectedToken(p.peek())
		}
	}

	return program
}

func (p *Parser) ParseVariableDeclaration() VariableDeclaration {
	endOfCursorLen := len(p.peek().Value)

	node := VariableDeclaration{
		BaseNode: BaseNode{
			Type:  "VariableDeclaration",
			Start: p.pos,
			End:   0,
		},
		Declarations: []ASTNode{},
		Kind:         p.peek().Value,
	}

	p.next()

	// expect identifier
	if p.peek().Type != IDENTIFIER {

		if p.peek().Value != "var" {
			oops.IllegalIdentifier(p.peek())
		}

		oops.ExpectedIdentifier(p.peek())
	}

	node.Declarations = append(node.Declarations, &VariableDeclarator{
		BaseNode: BaseNode{
			Type:  "VariableDeclarator",
			Start: p.pos + endOfCursorLen,
			End:   0,
		},
		Id: Identifier{
			Name: p.peek().Value,
			BaseNode: BaseNode{
				Type:  "Identifier",
				Start: p.pos + endOfCursorLen,
				End:   p.pos + endOfCursorLen + len(p.peek().Value),
			},
		},
		Init: nil,
	})

	endOfCursorLen += len(p.peek().Value)

	p.next()
	operator := p.peek().Value

	// expect assignment operator
	if operator != "=" {
		oops.ExpectedOperator(p.peek(), "=")
	}

	endOfCursorLen += len(p.peek().Value)
	p.next()

	node.Declarations[0].(*VariableDeclarator).Init = p.ParseAssignmentExpression(endOfCursorLen)
	endOfCursorLen += len(p.peek().Value)

	node.Declarations[0].(*VariableDeclarator).End = endOfCursorLen + p.pos
	node.BaseNode.End = endOfCursorLen + p.pos

	p.next()

	return node
}

func (p *Parser) ParseAssignmentExpression(endOfCursorLen int) ASTNode {
	switch p.peek().Type {
	case NUMBER:
		value, _ := strconv.Atoi(p.peek().Value)

		return Literal{
			BaseNode: BaseNode{
				Type:  "Literal",
				Start: p.pos + endOfCursorLen,
				End:   p.pos + len(p.peek().Value) + endOfCursorLen,
			},
			Value: value,
			Raw:   p.peek().Value,
		}

	case STRING:
		// TODO: Implement string parsing
		return nil
	case IDENTIFIER:
		// TODO: Implement identifier parsing
		return nil
	default:
		panic("Expected number, string or identifier")
	}
}
