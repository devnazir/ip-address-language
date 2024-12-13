package parser

import (
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseIdentifier(shouldCheckCallExpr ...interface{}) (ast.ASTNode, error) {
	v := utils.GetVariableName(p.peek().Value)
	rawValue := utils.GetVariableName(p.peek().RawValue)

	tree := ast.Identifier{
		Name: v,
		BaseNode: ast.BaseNode{
			Type:  ast.IdentifierTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Raw: rawValue,
	}

	p.next()

	if p.peek().Type == lx.TokenLeftBracket || p.peek().Type == lx.TokenDot {
		memberExpression, err := p.ParseMemberExpression(tree)
		if err != nil {
			return nil, err
		}

		return memberExpression, nil
	}

	if p.peek().Type == lx.TokenLeftParen && shouldCheckCallExpr != nil && len(shouldCheckCallExpr) > 0 {
		callExpression := p.parseCallExpression(tree)
		return callExpression, nil
	}

	return tree, nil
}
