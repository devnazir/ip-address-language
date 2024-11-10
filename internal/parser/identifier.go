package parser

import (
	"reflect"
	"strings"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseIdentifier() ast.Identifier {
	v, _ := utils.RemoveDoubleQuotes(p.peek().Value)
	trimmedName := strings.Trim(v, "$")

	ast := ast.Identifier{
		Name: strings.TrimSpace(trimmedName),
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
