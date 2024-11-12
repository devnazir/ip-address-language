package parser

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"

	lx "github.com/devnazir/gosh-script/internal/lexer"
)

func (p *Parser) ParseShellExpression() ast.ASTNode {
	switch p.peek().Value {
	case lx.KeywordSource:
		return p.ParseSourceDeclaration()
	case lx.KeywordEcho:
		return p.ParseEchoStatement()
	default:
		oops.UnexpectedKeywordError(p.peek())
	}

	return ast.ShellExpression{}
}
