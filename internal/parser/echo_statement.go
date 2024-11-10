package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseEchoStatement() ast.ASTNode {
	p.next()

	startStmtToken := p.peek()
	arguments := []ast.ASTNode{}
	flags := []string{}
	startLine := p.peek().Line

	for {
		if p.peek().Line != startLine || p.peek().Type == lx.TokenEOF {
			break
		}

		switch p.peek().Type {
		case lx.TokenString, lx.TokenIdentifier:
			arguments = append(arguments, p.ParseStringLiteral(&ParseStringLiteral{valueAsRaw: true}))
		case lx.TokenDollarSign:
			arguments = append(arguments, p.ParseIdentifier())
		case lx.TokenFlag:
			flags = append(flags, p.peek().Value)
			p.next()
		case lx.TokenNumber:
			arguments = append(arguments, p.ParseNumberLiteral())
		case lx.TokenIllegal:
			arguments = append(arguments, p.ParseIllegal())
		case lx.TokenSemicolon:
			p.next()
		default:
			arguments = append(arguments, p.ParseStringLiteral(nil))
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
