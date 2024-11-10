package parser

import (
	"reflect"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseIllegal() ast.Illegal {
	ast := ast.Illegal{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.Illegal{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: p.peek().Value,
		Raw:   p.peek().RawValue,
	}
	p.next()
	return ast
}
