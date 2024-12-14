package parser

import (
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"

	lx "github.com/devnazir/ip-address-language/internal/lexer"
)

func (p *Parser) ParseShellExpression() (ast.ASTNode, error) {
	switch p.peek().Value {
	case lx.KeywordSource:
		sourceDeclaration, err := p.ParseSourceDeclaration()
		if err != nil {
			panic(err)
		}

		return sourceDeclaration, nil
	case lx.KeywordEcho:
		echoStatement, err := p.ParseEchoStatement()
		if err != nil {
			panic(err)
		}
		return echoStatement, nil
	default:
		return nil, oops.SyntaxError(p.peek(), "Unknown shell expression")
	}
}
