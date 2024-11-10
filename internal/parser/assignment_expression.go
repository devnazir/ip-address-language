package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseAssignmentExpression(identToken lx.Token) ast.ASTNode {
	expression := p.EvaluateAssignmentExpression()

	return ast.AssignmentExpression{
		Identifier: ast.Identifier{
			Name: identToken.Value,
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.AssignmentExpression{}).Name(),
				Start: identToken.Start,
				End:   identToken.End,
				Line:  identToken.Line,
			},
		},
		Expression: expression,
	}
}

func (p *Parser) EvaluateAssignmentExpression() ast.ASTNode {
	var output []ast.ASTNode = []ast.ASTNode{}
	var operators []lx.Token = []lx.Token{}
	endLoop := false
	isBinaryExpression := false

	token := p.peek()
	startLine := token.Line

	for !endLoop && token.Type != lx.SHELL_KEYWORD {
		token := p.peek()

		if startLine != token.Line {
			// stop the loop
			endLoop = true
			break
		}

		switch token.Type {
		case NUMBER, STRING, IDENTIFIER:
			output = append(output, p.ParsePrimaryExpression())
		case OPERATOR:
			for len(operators) > 0 && Precedence[operators[len(operators)-1].Value] >= Precedence[token.Value] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}

			operators = append(operators, token)
			p.next()
		case LEFT_PAREN:
			operators = append(operators, token)
			p.next()
		case RIGHT_PAREN:
			for operators[len(operators)-1].Type != LEFT_PAREN {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
				continue
			}

			operators = operators[:len(operators)-1]
			p.next()
		default:
			endLoop = true
			break
		}
	}

	for len(operators) > 0 {
		isBinaryExpression = true
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	// utils.ParseToJson(output)

	if isBinaryExpression {
		return p.ParseBinaryExpression(output)
	}

	return output[0]
}

func (p *Parser) ParsePrimaryExpression() ast.ASTNode {
	switch p.peek().Type {
	case NUMBER:
		return p.ParseNumberLiteral()
	case STRING:
		return p.ParseStringLiteral(nil)
	case IDENTIFIER:
		return p.ParseIdentifier()
	default:
		panic("Expected a primary expression (number, string, or identifier)")
	}
}
