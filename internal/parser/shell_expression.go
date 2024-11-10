package parser

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"

	lx "github.com/devnazir/gosh-script/internal/lexer"
)

func (p *Parser) ParseShellExpression() ast.ASTNode {

	switch p.peek().Value {
	case lx.ECHO:
		return p.ParseEchoStatement()
	case lx.LS:
		return p.ParseLsStatement()
	default:
		oops.UnexpectedKeywordError(p.peek())
	}

	p.next()

	return ast.ShellExpression{}
}
