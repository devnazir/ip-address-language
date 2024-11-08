package parser

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"

	lx "github.com/devnazir/gosh-script/internal/lexer"
)

func (p *Parser) ParseShellExpression() ast.ASTNode {
	keyword := p.next()

	switch keyword.Value {
	case lx.ECHO:
		return p.ParseEchoStatement()
	case lx.LS:
		return p.ParseLsStatement()
	default:
		oops.UnexpectedKeywordError(keyword)
	}

	p.next()
	return ast.ShellExpression{}
}
