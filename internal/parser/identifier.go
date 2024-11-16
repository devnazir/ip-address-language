package parser

import (
	"reflect"
	"strings"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseIdentifier() ast.ASTNode {
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

	if p.peek().Type == lx.TokenLeftBracket || p.peek().Type == lx.TokenDot {
		return p.ParseMemberExpression(&ast)
	}

	return ast
}
