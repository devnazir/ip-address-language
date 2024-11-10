package parser

import (
	"reflect"

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
