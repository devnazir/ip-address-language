package parser

import (
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseEchoStatement() (ast.ShellExpression, error) {
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
			identifier, err := p.ParseIdentifier()
			if err != nil {
				return ast.ShellExpression{}, err
			}
			arguments = append(arguments, identifier)
		case lx.TokenFlag:
			flags = append(flags, p.peek().Value)
			p.next()
		case lx.TokenNumber:
			arguments = append(arguments, p.ParseNumberLiteral())
		case lx.TokenIllegal:
			arguments = append(arguments, p.ParseIllegal())
		case lx.TokenSubshell:
			arguments = append(arguments, p.ParseSubShell())
		case lx.TokenSemicolon:
			p.next()
		default:
			arguments = append(arguments, p.ParseStringLiteral(nil))
		}
	}

	tree := ast.ShellExpression{
		BaseNode: ast.BaseNode{
			Type:  ast.ShellExpressionTree,
			Start: startStmtToken.Start,
			End:   startStmtToken.End,
			Line:  startStmtToken.Line,
		},
		Expression: ast.EchoStatement{
			BaseNode: ast.BaseNode{
				Type:  ast.EchoStatementTree,
				Start: startStmtToken.Start,
				End:   startStmtToken.End,
				Line:  startStmtToken.Line,
			},
			Arguments: arguments,
			Flags:     flags,
		},
	}

	return tree, nil
}
