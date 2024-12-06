package parser

import (
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

func (p *Parser) ParseStringTemplateLiteral() ast.StringTemplateLiteral {
	removedQuotes := p.peek().Value[1 : len(p.peek().Value)-1]
	newLexer := lx.NewLexer(removedQuotes, "")
	tokens := newLexer.Tokenize()

	parts := []ast.ASTNode{}

	parseAsStringLiteral := func(token lx.Token) ast.ASTNode {
		return ast.StringLiteral{
			Value: token.Value,
			BaseNode: ast.BaseNode{
				Type:  ast.StringLiteralTree,
				Start: token.Start + len(token.Value) + len(*tokens),
				End:   token.End + len(*tokens) + len(token.Value),
				Line:  token.Line,
			},
			Raw: token.RawValue,
		}
	}

	for _, token := range *tokens {

		if token.Value == "" {
			continue
		}

		switch token.Type {
		case lx.TokenIdentifier:
			parts = append(parts, parseAsStringLiteral(token))
		case lx.TokenSubshell:
			matcherArgs := utils.FindSubShellArgs(token.Value)
			parts = append(parts, ast.SubShell{
				Arguments: matcherArgs[1],
				BaseNode: ast.BaseNode{
					Type:  ast.SubShellTree,
					Start: token.Start + len(token.Value) + len(*tokens),
					End:   token.End + len(token.Value) + len(*tokens),
					Line:  token.Line,
				},
			})
		case lx.TokenDollarSign:
			name := token.Value[1:]
			parts = append(parts, ast.Identifier{
				Name: name,
				BaseNode: ast.BaseNode{
					Type:  ast.IdentifierTree,
					Start: token.Start + len(token.Value) + len(*tokens), // +1 to remove the dollar sign
					End:   token.End + len(token.Value) + len(*tokens),
					Line:  token.Line,
				},
			})
		default:
			parts = append(parts, parseAsStringLiteral(token))
		}
	}

	tree := ast.StringTemplateLiteral{
		Parts: parts,
		BaseNode: ast.BaseNode{
			Type:  ast.StringTemplateLiteralTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
	}

	p.next()
	return tree
}
