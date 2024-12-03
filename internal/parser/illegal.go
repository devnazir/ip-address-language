package parser

import (
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseIllegal() ast.Illegal {
	tree := ast.Illegal{
		BaseNode: ast.BaseNode{
			Type:  ast.IllegalTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: p.peek().Value,
		Raw:   p.peek().RawValue,
	}
	p.next()
	return tree
}
