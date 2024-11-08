package parser

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseLiteral() ast.Literal {
	value := p.peek().Value
	var literalValue interface{}

	switch p.peek().Type {
	case NUMBER:
		if strings.Contains(value, ".") {
			literalValue, _ = strconv.ParseFloat(value, 64)
		} else {
			literalValue, _ = strconv.Atoi(value)
		}
	default:
		literalValue = value
	}

	ast := ast.Literal{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.Literal{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Value: literalValue,
		Raw:   value,
	}
	p.next()
	return ast
}
