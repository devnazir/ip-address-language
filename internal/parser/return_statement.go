package parser

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
)

type ParseReturnStatementParams struct {
	isNextReturn bool
}

func (p *Parser) ParseReturnStatement(params ParseReturnStatementParams) (ast.ReturnStatement, error) {
	token := p.peek()

	if !params.isNextReturn {
		if !p.TokenValueIs(lx.KeywordReturn) {
			return ast.ReturnStatement{}, oops.RuntimeError(token, "Expect 'return' keyword")
		}

		p.next()
	}

	if p.peek().Type == lx.TokenRightCurly {
		return ast.ReturnStatement{
			BaseNode: ast.BaseNode{
				Type:  ast.ReturnStatementTree,
				Start: token.Start,
				End:   token.End,
				Line:  token.Line,
			},
			Arguments: []ast.ASTNode{},
		}, nil
	}

	expr, err := p.ParsePrimaryExpression()

	if err != nil {
		return ast.ReturnStatement{}, err
	}

	arguments := []ast.ASTNode{
		expr,
	}

	if p.TokenTypeIs(lx.TokenComma) {
		p.next()
		nextReturn, err := p.ParseReturnStatement(ParseReturnStatementParams{
			isNextReturn: true,
		})
		if err != nil {
			return ast.ReturnStatement{}, err
		}
		arguments = append(arguments, nextReturn.Arguments...)
	}

	return ast.ReturnStatement{
		BaseNode: ast.BaseNode{
			Type:  ast.ReturnStatementTree,
			Start: token.Start,
			End:   token.End,
			Line:  token.Line,
		},
		Arguments: arguments,
	}, nil
}
