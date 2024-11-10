package parser

import (
	"reflect"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseNumberLiteral() ast.NumberLiteral {
	value := p.peek().Value

	numVal, _ := utils.InferType(value)

	ast := ast.NumberLiteral{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.NumberLiteral{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: numVal,
		Raw:   value,
	}
	p.next()
	return ast
}
