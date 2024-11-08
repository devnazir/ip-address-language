package parser

import (
	"reflect"
	"strings"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseIdentifier() ast.Identifier {
	trimmedName := strings.Trim(p.peek().Value, "$;")

	ast := ast.Identifier{
		Name: trimmedName,
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.Identifier{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
	}
	p.next()
	return ast
}
