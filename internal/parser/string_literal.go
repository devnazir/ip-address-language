package parser

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
	"github.com/devnazir/ip-address-language/pkg/utils"
)

type ParseStringLiteral struct {
	valueAsRaw bool
}

func (p *Parser) ParseStringLiteral(params *ParseStringLiteral) ast.StringLiteral {
	start := p.peek().Start
	hasDoubleQuote := p.TokenTypeIs(lx.TokenDoubleQuote)

	if p.TokenTypeIs(lx.TokenDoubleQuote) {
		// skip double quote
		p.next()
	}

	value := ""
	rawValue := ""

	for p.peek().Type != lx.TokenDoubleQuote && hasDoubleQuote {
		value += p.peek().Value
		rawValue += p.peek().RawValue
		p.next()
	}

	if !hasDoubleQuote {
		value = p.peek().Value
		rawValue = p.peek().RawValue
	}

	if params != nil && params.valueAsRaw {
		value = rawValue
	}

	finalValue, _ := utils.RemoveDoubleQuotes(value)

	ast := ast.StringLiteral{
		BaseNode: ast.BaseNode{
			Type:  ast.StringLiteralTree,
			Start: start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: finalValue,
		Raw:   rawValue,
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
