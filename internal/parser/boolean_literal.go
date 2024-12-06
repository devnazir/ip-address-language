package parser

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseBooleanLiteral() ast.BooleanLiteral {
	boolVal, _ := utils.InferType(p.peek().Value)

	tree := ast.BooleanLiteral{
		BaseNode: ast.BaseNode{
			Type:  ast.BooleanLiteralTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: boolVal.(bool),
		Raw:   p.peek().Value,
	}
	p.next()
	return tree
}
