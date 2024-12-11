package parser

import (
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseMemberExpression(ident ast.Identifier) (ast.MemberExpression, error) {
	token := p.next()
	property, err := p.ParsePrimaryExpression()
	if err != nil {
		return ast.MemberExpression{}, err
	}

	base := ast.MemberExpression{
		BaseNode: ast.BaseNode{
			Type:  ast.MemberExpressionTree,
			Start: ident.Start,
			End:   token.End,
			Line:  ident.Line,
		},
		Object:   ident,
		Property: property,
		Computed: false,
	}

	if token.Type == lx.TokenLeftBracket {
		base.Computed = true
	}

	for {
		switch p.peek().Type {
		case lx.TokenDot:
			p.next()
			property, err := p.ParseIdentifier()
			if err != nil {
				return ast.MemberExpression{}, err
			}

			base = ast.MemberExpression{
				BaseNode: ast.BaseNode{
					Type:  ast.MemberExpressionTree,
					Start: token.Start,
					End:   token.End,
					Line:  p.peek().Line,
				},
				Object:   base,
				Property: property,
				Computed: false,
			}

		case lx.TokenLeftBracket:
			p.next()
			property, err := p.ParsePrimaryExpression()
			if err != nil {
				return ast.MemberExpression{}, err
			}

			if p.peek().Type != lx.TokenRightBracket {
				return ast.MemberExpression{}, oops.SyntaxError(p.peek(), "Expected ']' after computed property")
			}

			p.next()

			base = ast.MemberExpression{
				BaseNode: ast.BaseNode{
					Type:  ast.MemberExpressionTree,
					Start: token.Start,
					End:   token.End,
					Line:  p.peek().Line,
				},
				Object:   base,
				Property: property,
				Computed: true,
			}

		case lx.TokenRightBracket:
			p.next()

		case lx.TokenLeftParen:
			ident := ast.Identifier{
				BaseNode: ast.BaseNode{
					Type:  ast.IdentifierTree,
					Start: base.Start,
					End:   base.End,
					Line:  base.Line,
				},
				Name: base.Property.(ast.Identifier).Name,
			}

			callExpression := p.parseCallExpression(ident)
			base.Property = callExpression

		case lx.TokenRightParen:
			p.next()

		default:
			return base, nil
		}
	}
}
