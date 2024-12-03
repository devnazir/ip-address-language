package parser

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseNumberLiteral() ast.NumberLiteral {
	value := p.peek().Value

	numVal, _ := utils.InferType(value)

	tree := ast.NumberLiteral{
		BaseNode: ast.BaseNode{
			Type:  ast.NumberLiteralTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: numVal,
		Raw:   value,
	}
	p.next()
	return tree
}
