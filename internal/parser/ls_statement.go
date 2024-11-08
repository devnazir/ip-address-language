package parser

import (
	"reflect"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseLsStatement() ast.ASTNode {
	startStmtToken := p.peek()

	return ast.ShellExpression{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.ShellExpression{}).Name(),
			Start: startStmtToken.Start,
			End:   startStmtToken.End,
			Line:  startStmtToken.Line,
		},
		Expression: ast.LsStatement{
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.LsStatement{}).Name(),
				Start: startStmtToken.Start,
				End:   startStmtToken.End,
				Line:  startStmtToken.Line,
			},
		},
	}
}
