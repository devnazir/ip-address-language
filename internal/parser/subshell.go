package parser

import (
	"fmt"
	"reflect"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (p *Parser) ParseSubShell() ast.ASTNode {
	matcherArgs := utils.FindSubShellArgs(p.peek().Value)

	if len(matcherArgs) == 0 {
		fmt.Println("Invalid subshell expression")
		return nil
	}

	ast := ast.SubShell{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.SubShell{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Arguments: matcherArgs[1],
	}

	p.next()
	return ast
}
