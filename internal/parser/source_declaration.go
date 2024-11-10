package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseSourceDeclaration() ast.ASTNode {
	token := p.next()
	sources := []ast.ASTNode{}

	switch p.peek().Type {
	case lx.TokenString:
		sources = append(sources, p.ParseStringLiteral(nil))

	case lx.TokenLeftParen:
		p.next()
		endLoop := false

		for !endLoop {
			switch p.peek().Type {
			case lx.TokenString:
				sources = append(sources, p.ParseStringLiteral(nil))
			case lx.TokenRightParen:
				endLoop = true
				p.next()
			default:
				oops.ExpectedTokenError(p.peek(), ")")
				p.next()
			}
		}
	default:
		oops.UnexpectedTokenError(p.peek(), "")
	}

	return ast.SourceDeclaration{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.SourceDeclaration{}).Name(),
			Start: token.Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Sources: sources,
	}
}
