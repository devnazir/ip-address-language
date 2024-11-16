package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseFunctionDeclaration() ast.ASTNode {
	node := ast.FunctionDeclaration{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.FunctionDeclaration{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Identifier: ast.Identifier{},
		Body:       []ast.ASTNode{},
		Parameters: []ast.Identifier{},
	}

	p.next()
	// consume the function name
	fnName := p.next().Value

	node.Identifier = ast.Identifier{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.Identifier{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Name: fnName,
	}

	// consume the left parenthesis
	leftParen := p.next()

	if leftParen.Type != lx.TokenLeftParen {
		panic("Expected left parenthesis")
	}
	for p.peek().Type != lx.TokenRightParen && p.peek().Type != lx.TokenLeftCurly {

		isRestParameter := false
		dotLen := 0
		maxDotLen := 3

		for p.peek().Type == lx.TokenDot {
			dotLen++

			if dotLen > maxDotLen {
				panic("Expected identifier")
			}

			p.next()
		}

		if dotLen == maxDotLen {
			isRestParameter = true
		}

		if p.peek().Type != lx.TokenIdentifier {
			panic("Expected identifier")
		}

		if p.peek().Type == lx.TokenIdentifier {

			ident := ast.Identifier{
				BaseNode: ast.BaseNode{
					Type:  reflect.TypeOf(ast.Identifier{}).Name(),
					Start: p.peek().Start,
					End:   p.peek().End,
					Line:  p.peek().Line,
				},
				Name: p.peek().Value,
			}

			if isRestParameter {
				ident.IsRestParameter = true
			}

			node.Parameters = append(node.Parameters, ident)
			p.next()

			if p.peek().Type == lx.TokenComma {
				p.next()
				continue
			}

			if p.peek().Type == lx.TokenRightParen {
				break
			} else {
				panic("Expected comma or right parenthesis")
			}
		}

	}

	// consume the right parenthesis
	rightParen := p.next()

	if rightParen.Type != lx.TokenRightParen {
		panic("Expected right parenthesis")
	}

	// consume the left curly brace
	leftCurly := p.next()

	if leftCurly.Type != lx.TokenLeftCurly {
		panic("Expected left curly brace")
	}

	body := []ast.ASTNode{}
	for p.peek().Type != lx.TokenRightCurly {
		program := p.ParseBodyProgram(&ast.Program{
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.Program{}).Name(),
				Start: p.peek().Start,
				End:   0,
				Line:  p.peek().Line,
			},
			Body: node.Body,
		})

		body = append(body, program.(*ast.Program).Body[0])
	}

	node.Body = body

	if p.peek().Type != lx.TokenRightCurly {
		panic("Expected right curly brace")
	}

	return node
}
