package parser

import (
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/utils"
)

type ParseStringLiteral struct {
	valueAsRaw bool
}

func (p *Parser) ParseStringLiteral(params *ParseStringLiteral) ast.StringLiteral {
	value := p.peek().Value

	if params != nil {
		if params.valueAsRaw {
			value = p.peek().RawValue
		}
	}

	v, _ := utils.RemoveDoubleQuotes(value)

	ast := ast.StringLiteral{
		BaseNode: ast.BaseNode{
			Type:  ast.StringLiteralTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: v,
		Raw:   p.peek().RawValue,
	}
	p.next()
	return ast
}

func (p *Parser) ParseStringTemplateLiteral() (ast.StringTemplateLiteral, error) {
	parts := []ast.ASTNode{}
	start := p.peek().Start

	if p.peek().Type == lx.TokenTickQuote {
		p.next()

		for p.peek().Type != lx.TokenTickQuote {
			if p.peek().Type == lx.TokenEOF {
				break
			}

			switch p.peek().Type {
			case lx.TokenIdentifier:
				parts = append(parts, p.ParseStringLiteral(&ParseStringLiteral{valueAsRaw: true}))
			case lx.TokenDollarSign:
				identifier, err := p.ParseIdentifier()
				if err != nil {
					panic(err)
				}
				parts = append(parts, identifier)
			case lx.TokenSubshell:
				parts = append(parts, p.ParseSubShell())
			default:
				if p.peek().Type == lx.TokenRightCurly {
					p.next()
					continue
				}
				parts = append(parts, p.ParseStringLiteral(nil))
			}
		}

		if p.peek().Type == lx.TokenTickQuote {
			p.next()
			return ast.StringTemplateLiteral{
				Parts: parts,
				BaseNode: ast.BaseNode{
					Type:  ast.StringTemplateLiteralTree,
					Start: start,
					End:   p.peek().End,
					Line:  p.peek().Line,
				},
			}, nil
		}
	}

	return ast.StringTemplateLiteral{}, oops.SyntaxError(p.peek(), "Invalid string template literal")
}
