package parser

import (
	"strings"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseIdentifier() (ast.ASTNode, error) {
	v, _ := utils.RemoveDoubleQuotes(p.peek().Value)
	trimmedName := strings.Trim(v, "$")

	tree := ast.Identifier{
		Name: strings.TrimSpace(trimmedName),
		BaseNode: ast.BaseNode{
			Type:  ast.IdentifierTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
	}

	p.next()

	if p.peek().Type == lx.TokenLeftBracket || p.peek().Type == lx.TokenDot {
		memberExpression, err := p.ParseMemberExpression(tree)
		if err != nil {
			return nil, err
		}

		return memberExpression, nil
	}

	return tree, nil
}
