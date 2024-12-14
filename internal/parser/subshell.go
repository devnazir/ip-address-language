package parser

import (
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/utils"
)

func (p *Parser) ParseSubShell() ast.ASTNode {
	matcherArgs := utils.FindSubShellArgs(p.peek().Value)

	if len(matcherArgs) == 0 {
		return nil
	}

	ast := ast.SubShell{
		BaseNode: ast.BaseNode{
			Type:  ast.SubShellTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Arguments: matcherArgs[1],
	}

	p.next()
	return ast
}
