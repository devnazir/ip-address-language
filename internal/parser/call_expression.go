package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) parseCallExpression(ident *ast.Identifier) ast.CallExpression {
	p.next()
	arguments := []ast.ASTNode{}

	for p.peek().Type != lx.TokenRightParen {

		if p.peek().Type == lx.TokenComma {
			p.next()
			continue
		}

		arguments = append(arguments, p.ParsePrimaryExpression())
	}

	p.next()

	return ast.CallExpression{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.CallExpression{}).Name(),
			Start: ident.Start,
			End:   p.peek().End,
			Line:  ident.Line,
		},
		Callee: ast.Identifier{
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.Identifier{}).Name(),
				Start: ident.Start,
				End:   ident.End,
				Line:  ident.Line,
			},
			Name: ident.Name,
		},
		Arguments: arguments,
	}
}
