package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) generateAlias(sources *[]ast.Source) {
	identValue := p.peek().Value

	if identValue != "as" {
		oops.UnexpectedTokenError(p.peek(), "as")
	}

	p.next()

	identValue = p.peek().Value

	(*sources)[len(*sources)-1] = ast.Source{
		StringLiteral: (*sources)[len(*sources)-1].StringLiteral,
		Alias:         identValue,
	}

	p.next()
}

func (p *Parser) ParseSourceDeclaration() ast.ASTNode {
	token := p.next()
	sources := &[]ast.Source{}

	switch p.peek().Type {
	case lx.TokenString:
		*sources = append(*sources, p.ParseSource(""))

		if p.peek().Type == lx.TokenIdentifier {
			p.generateAlias(sources)
		}

	case lx.TokenLeftParen:
		p.next()
		endLoop := false

		for !endLoop {
			switch p.peek().Type {
			case lx.TokenString:
				*sources = append(*sources, p.ParseSource(""))

			case lx.TokenIdentifier:
				p.generateAlias(sources)

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
		Sources: *sources,
	}
}

func (p *Parser) ParseSource(alias string) ast.Source {
	ast := ast.Source{
		StringLiteral: p.ParseStringLiteral(nil),
	}

	if alias != "" {
		ast.Alias = alias
	}

	return ast
}
