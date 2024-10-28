package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

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
			if p.peek().Value == lx.VAR {
				program.Body = append(program.Body, p.ParseVariableDeclaration())
			} else {
				oops.UnexpectedKeyword(p.peek())
			}
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

func (p *Parser) ParseVariableDeclaration() VariableDeclaration {

	node := VariableDeclaration{
		BaseNode: BaseNode{
			Type:  "VariableDeclaration",
			Start: p.pos,
			End:   0,
		},
		Declarations: []ASTNode{},
		Kind:         p.peek().Value,
	}

	p.next() // skip "var"

	// expect identifier
	if p.peek().Type != IDENTIFIER {

		if p.peek().Value != lx.VAR {
			oops.IllegalIdentifier(p.peek())
		}

		oops.ExpectedIdentifier(p.peek())
	}

	node.Declarations = append(node.Declarations, &VariableDeclarator{
		BaseNode: BaseNode{
			Type:  "VariableDeclarator",
			Start: p.peek().Start,
			End:   0,
		},
		Id: Identifier{
			Name: p.peek().Value,
			BaseNode: BaseNode{
				Type:  "Identifier",
				Start: p.peek().Start,
				End:   p.peek().End,
			},
		},
		Init: nil,
	})

	p.next() // skip identifier, next to assignment operator

	// check if the next token has primitive type
	if p.peek().Type == PRIMITIVE_TYPE {
		primitiveType := p.peek().Value
		node.TypeAnnotation = primitiveType
		p.next()
	}

	// expect assignment operator
	operator := p.peek().Value
	if operator != "=" {
		oops.UnexpectedToken(p.peek(), "=")
	}

	p.next() // next to assignment expression

	node.Declarations[0].(*VariableDeclarator).Init = p.ParseAssignmentExpression()
	node.Declarations[0].(*VariableDeclarator).End = p.peek().End
	node.BaseNode.End = p.peek().End

	p.next()

	if node.TypeAnnotation != "" {
		valueType := reflect.TypeOf(node.Declarations[0].(*VariableDeclarator).Init.(Literal).Value).String()

		if valueType != node.TypeAnnotation {
			oops.TypeMismatch(p.peek(), node.TypeAnnotation, valueType)
		}

	} else {
		// infer type
		valueType := reflect.TypeOf(node.Declarations[0].(*VariableDeclarator).Init.(Literal).Value).String()
		node.TypeAnnotation = valueType
	}

	return node
}

func (p *Parser) ParseAssignmentExpression() ASTNode {
	switch p.peek().Type {
	case NUMBER:

		hasDecimal := strings.Contains(p.peek().Value, ".")
		var value interface{} = p.peek().Value

		// Currently, only support int and float64
		if hasDecimal {
			value, _ = strconv.ParseFloat(p.peek().Value, 64)
		} else {
			value, _ = strconv.Atoi(p.peek().Value)
		}

		return Literal{
			BaseNode: BaseNode{
				Type:  "Literal",
				Start: p.peek().Start,
				End:   p.peek().End,
			},
			Value: value,
			Raw:   p.peek().Value,
		}

	case STRING:
		return Literal{
			BaseNode: BaseNode{
				Type:  "Literal",
				Start: p.peek().Start,
				End:   p.peek().End,
			},
			Value: p.peek().Value,
			Raw:   p.peek().Value,
		}
	case IDENTIFIER:
		// TODO: Implement identifier parsing
		return nil
	default:
		panic("Expected number, string or identifier")
	}
}
