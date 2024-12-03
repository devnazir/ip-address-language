package parser

import (
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) parseCallExpression(ident ast.Identifier) ast.CallExpression {
	p.next()
	arguments := []ast.ASTNode{}

	for p.peek().Type != lx.TokenRightParen {

		if p.peek().Type == lx.TokenComma {
			p.next()
			continue
		}

		primaryExpression, err := p.ParsePrimaryExpression()
		if err != nil {
			panic(err)
		}

		arguments = append(arguments, primaryExpression)
	}

	p.next()

	return ast.CallExpression{
		BaseNode: ast.BaseNode{
			Type:  ast.CallExpressionTree,
			Start: ident.Start,
			End:   p.peek().End,
			Line:  ident.Line,
		},
		Callee: ast.Identifier{
			BaseNode: ast.BaseNode{
				Type:  ast.IdentifierTree,
				Start: ident.Start,
				End:   ident.End,
				Line:  ident.Line,
			},
			Name: ident.Name,
		},
		Arguments: arguments,
	}
}
