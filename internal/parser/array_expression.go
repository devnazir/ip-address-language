package parser

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
)

func (p *Parser) ParseArrayExpression() (ast.ASTNode, error) {
	var elements []ast.ASTNode = []ast.ASTNode{}
	token := p.peek()

	if token.Type == lx.TokenLeftBracket {
		p.next()

		for p.peek().Type != lx.TokenRightBracket {
			if p.peek().Type == lx.TokenComma {
				p.next()
				continue
			}

			if p.peek().Type == lx.TokenEOF {
				break
			}

			expr, err := p.ParsePrimaryExpression()

			if err != nil {
				panic(err)
			}

			elements = append(elements, expr)
		}

		if p.peek().Type == lx.TokenRightBracket {
			p.next()
			return ast.ArrayExpression{
				BaseNode: ast.BaseNode{
					Type:  ast.ArrayExpressionTree,
					Start: token.Start,
					End:   token.End,
					Line:  token.Line,
				},
				Elements: elements,
			}, nil
		}
	}

	return nil, oops.SyntaxError(p.peek(), "Expected closing bracket ']'")
}
