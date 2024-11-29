package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
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
			Type:  reflect.TypeOf(ast.StringLiteral{}).Name(),
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

func (p *Parser) ParseStringTemplateLiteral() ast.StringTemplateLiteral {
	removedQuotes := p.peek().Value[1 : len(p.peek().Value)-1]
	newLexer := lx.NewLexer(removedQuotes, "")
	tokens := newLexer.Tokenize()

	parts := []ast.ASTNode{}

	parseAsStringLiteral := func(token lx.Token) ast.ASTNode {
		return ast.StringLiteral{
			Value: token.Value,
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.StringLiteral{}).Name(),
				Start: token.Start,
				End:   token.End,
				Line:  token.Line,
			},
			Raw: token.RawValue,
		}
	}

	for _, token := range *tokens {
		switch token.Type {
		case lx.TokenIdentifier:
			parts = append(parts, parseAsStringLiteral(token))
		case lx.TokenDollarSign:
			name := token.Value[1:]
			parts = append(parts, ast.Identifier{
				Name: name,
				BaseNode: ast.BaseNode{
					Type:  reflect.TypeOf(ast.Identifier{}).Name(),
					Start: token.Start,
					End:   token.End,
					Line:  token.Line,
				},
			})
		default:
			parts = append(parts, parseAsStringLiteral(token))
		}
	}

	ast := ast.StringTemplateLiteral{
		Parts: parts,
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.StringTemplateLiteral{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
	}

	p.next()
	return ast
}
