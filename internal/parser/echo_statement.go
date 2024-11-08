package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseEchoStatement() ast.ASTNode {
	startStmtToken := p.peek()
	arguments := []ast.ASTNode{}
	flags := []string{}

	for p.peek().Type != lx.SEMICOLON && p.peek().Type != lx.EOF {
		switch p.peek().Type {
		case STRING, lx.IDENTIFIER:
			arguments = append(arguments, p.ParseLiteral())
		case lx.DOLLAR_SIGN:
			arguments = append(arguments, p.ParseIdentifier())
		case lx.FLAG:
			flags = append(flags, p.peek().Value)
			p.next()
		default:
			arguments = append(arguments, p.ParseLiteral())
		}
	}

	return ast.ShellExpression{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.ShellExpression{}).Name(),
			Start: startStmtToken.Start,
			End:   startStmtToken.End,
			Line:  startStmtToken.Line,
		},
		Expression: ast.EchoStatement{
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.EchoStatement{}).Name(),
				Start: startStmtToken.Start,
				End:   startStmtToken.End,
				Line:  startStmtToken.Line,
			},
			Arguments: arguments,
			Flags:     flags,
		},
	}
}
