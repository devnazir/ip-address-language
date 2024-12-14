package parser

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
	"github.com/devnazir/ip-address-language/pkg/utils"
)

func (p *Parser) ParseIfStatement() (ast.IfStatement, error) {
	if !p.TokenValueIs(lx.KeywordIf) {
		return ast.IfStatement{}, oops.SyntaxError(p.peek(), "expected if")
	}

	var output []ast.ASTNode = []ast.ASTNode{}
	var operators []lx.Token = []lx.Token{}
	ifStatement := ast.IfStatement{
		BaseNode: ast.BaseNode{
			Type:  ast.IFStatementTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
	}

	p.next()

	for p.peek().Type != lx.TokenLeftCurly {
		switch p.peek().Type {
		case
			lx.TokenNumber,
			lx.TokenDoubleQuote,
			lx.TokenIdentifier,
			lx.TokenDollarSign,
			lx.TokenTickQuote,
			lx.TokenBoolean:

			primaryExpression, err := p.ParsePrimaryExpression()
			if err != nil {
				panic(err)
			}

			output = append(output, primaryExpression)

		case lx.TokenOperator:

			for len(operators) > 0 && ComparisonPrecedence[operators[len(operators)-1].Value] >= ComparisonPrecedence[p.peek().Value] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}

			operators = append(operators, p.peek())
			p.next()

		case lx.TokenLeftParen:
			operators = append(operators, p.peek())
			p.next()
		case lx.TokenRightParen:
			for operators[len(operators)-1].Type != lx.TokenLeftParen {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
				continue
			}

			operators = operators[:len(operators)-1]
			p.next()

		default:
			utils.PrintJson(p.peek())
			return ifStatement, oops.SyntaxError(p.peek(), "unexpected token")
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	expr := p.ParseBinaryExpression(output)

	if !p.TokenTypeIs(lx.TokenLeftCurly) {
		return ifStatement, oops.SyntaxError(p.peek(), "expected {")
	}

	p.next()

	body := ast.BodyProgram{}
	for p.peek().Type != lx.TokenRightCurly {
		program, err := p.ParseBodyProgram()

		if err != nil {
			return ast.IfStatement{}, err
		}

		body = append(body, program...)
	}

	ifStatement.Condition = expr
	ifStatement.Consequent = body

	if !p.TokenTypeIs(lx.TokenRightCurly) {
		return ifStatement, oops.SyntaxError(p.peek(), "expected }")
	}

	p.next()

	if p.TokenValueIs(lx.KeywordElse) {
		p.next()

		if p.TokenValueIs(lx.KeywordIf) {
			elseIfStatement, err := p.ParseIfStatement()
			if err != nil {
				panic(err)
			}

			ifStatement.Alternate = elseIfStatement
		} else {
			if !p.TokenTypeIs(lx.TokenLeftCurly) {
				return ifStatement, oops.SyntaxError(p.peek(), "expected {")
			}

			p.next()

			body := ast.BodyProgram{}

			for p.peek().Type != lx.TokenRightCurly && p.peek().Value != lx.KeywordIf {
				program, err := p.ParseBodyProgram()

				if err != nil {
					return ast.IfStatement{}, err
				}

				body = append(body, program...)
			}

			ifStatement.Alternate = body

			if !p.TokenTypeIs(lx.TokenRightCurly) {
				return ifStatement, oops.SyntaxError(p.peek(), "expected }")
			}

			p.next()
		}
	}

	return ifStatement, nil
}
