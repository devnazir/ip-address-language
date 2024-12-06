package parser

import (
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseObjectExpression() (ast.ObjectExpression, error) {
	properties := []ast.Property{}

	for p.peek().Type != lx.TokenRightCurly {
		p.next()
		if p.peek().Type == lx.TokenComma {
			p.next()
			continue
		}

		if p.peek().Type == lx.TokenEOF {
			break
		}

		key, keyType := utils.InferType(p.peek().Value)

		if keyType != "string" {
			return ast.ObjectExpression{}, oops.SyntaxError(p.peek(), "Expected identifier")
		}

		p.next()

		if p.peek().Type != lx.TokenColon {
			return ast.ObjectExpression{}, oops.SyntaxError(p.peek(), "Expected colon ':'")
		}

		p.next()

		value, err := p.ParsePrimaryExpression()
		if err != nil {
			return ast.ObjectExpression{}, err
		}

		properties = append(properties, ast.Property{
			Key:   key.(string),
			Value: value,
		})

	}

	if p.peek().Type == lx.TokenRightCurly {
		p.next()
		return ast.ObjectExpression{
			BaseNode: ast.BaseNode{
				Type:  ast.ObjectExpressionTree,
				Start: p.peek().Start,
				End:   p.peek().End,
				Line:  p.peek().Line,
			},
			Properties: properties,
		}, nil
	}

	return ast.ObjectExpression{}, oops.SyntaxError(p.peek(), "Expected closing curly brace '}'")
}
