package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseMemberExpression(ident *ast.Identifier) ast.ASTNode {
	token := p.next()
	property := p.ParsePrimaryExpression()

	base := ast.MemberExpression{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.MemberExpression{}).Name(),
			Start: ident.Start,
			End:   token.End,
			Line:  ident.Line,
		},
		Object:   *ident,
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
			property := p.ParseIdentifier()

			base = ast.MemberExpression{
				BaseNode: ast.BaseNode{
					Type:  reflect.TypeOf(ast.MemberExpression{}).Name(),
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
			property := p.ParsePrimaryExpression()

			if p.peek().Type != lx.TokenRightBracket {
				panic("Expected ']' after computed property")
			}

			p.next()

			base = ast.MemberExpression{
				BaseNode: ast.BaseNode{
					Type:  reflect.TypeOf(ast.MemberExpression{}).Name(),
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

		default:
			return base
		}
	}
}
