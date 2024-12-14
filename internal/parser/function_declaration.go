package parser

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
)

func (p *Parser) ParseFunctionDeclaration() (ast.FunctionDeclaration, error) {
	node := ast.FunctionDeclaration{
		BaseNode: ast.BaseNode{
			Type:  ast.FunctionDeclarationTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Identifier: ast.Identifier{},
		Body:       []ast.ASTNode{},
		Parameters: []ast.Identifier{},
	}

	isAnonymousFn := false

	p.next()
	// consume the function name
	identToken := p.next()

	if identToken.Type == lx.TokenLeftParen {
		isAnonymousFn = true
	}

	node.IsAnonymous = isAnonymousFn

	if !isAnonymousFn {
		node.Identifier = ast.Identifier{
			BaseNode: ast.BaseNode{
				Type:  ast.IdentifierTree,
				Start: p.peek().Start,
				End:   p.peek().End,
				Line:  p.peek().Line,
			},
			Name: identToken.Value,
		}
	}

	// consume the left parenthesis
	if !isAnonymousFn {
		nextToken := p.next()

		if nextToken.Type != lx.TokenLeftParen && !isAnonymousFn {
			return ast.FunctionDeclaration{}, oops.SyntaxError(nextToken, "Expected left parenthesis")
		}
	}

	for p.peek().Type != lx.TokenRightParen && p.peek().Type != lx.TokenLeftCurly {
		isRestParameter := false
		dotLen := 0
		maxDotLen := 3

		for p.peek().Type == lx.TokenDot {
			dotLen++

			if dotLen > maxDotLen {
				return ast.FunctionDeclaration{}, oops.SyntaxError(p.peek(), "Invalid rest parameter")
			}

			p.next()
		}

		if dotLen == maxDotLen {
			isRestParameter = true
		}

		if p.peek().Type != lx.TokenIdentifier {
			return ast.FunctionDeclaration{}, oops.SyntaxError(p.peek(), "Expected identifier")
		}

		if p.peek().Type == lx.TokenIdentifier {

			ident := ast.Identifier{
				BaseNode: ast.BaseNode{
					Type:  ast.IdentifierTree,
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
				return ast.FunctionDeclaration{}, oops.SyntaxError(p.peek(), "Expected comma or right parenthesis")
			}
		}
	}

	rightParenToken := p.next()

	if rightParenToken.Type != lx.TokenRightParen {
		return ast.FunctionDeclaration{}, oops.SyntaxError(rightParenToken, "Expected right parenthesis")
	}

	// consume the left curly brace
	leftCurly := p.next()

	if leftCurly.Type != lx.TokenLeftCurly {
		return ast.FunctionDeclaration{}, oops.SyntaxError(leftCurly, "Expected left curly brace")
	}

	body := []ast.ASTNode{}
	for p.peek().Type != lx.TokenRightCurly {
		program, err := p.ParseBodyProgram()

		if err != nil {
			return ast.FunctionDeclaration{}, err
		}

		body = append(body, program...)
	}

	node.Body = body

	if p.peek().Type != lx.TokenRightCurly {
		return ast.FunctionDeclaration{}, oops.SyntaxError(p.peek(), "Expected right curly brace")
	}

	p.next()
	return node, nil
}
