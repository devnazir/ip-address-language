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
	var output []ast.ASTNode
	var operators []lx.Token
	endLoop := false

	for !endLoop {
		token := p.peek()

		switch token.Type {
		case NUMBER, STRING:
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
		case IDENTIFIER:
			output = append(output, p.ParsePrimaryExpression())
		case lx.SHELL_KEYWORD:
			output = append(output, p.ParseShellExpression())
		default:
			endLoop = true
			break
		}
	}

	for len(operators) > 0 {
		output = append(output, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}

	isBinaryExpression := false
	reflectLastOutput := reflect.TypeOf(output[len(output)-1])

	if reflectLastOutput == reflect.TypeOf(lx.Token{}) {
		if lx.TokenType(output[len(output)-1].(lx.Token).Type) == OPERATOR {
			isBinaryExpression = true
		}
	}

	if isBinaryExpression {
		return p.ParseBinaryExpression(output)
	}

	return output[0]
}

func (p *Parser) ParsePrimaryExpression() ast.ASTNode {
	switch p.peek().Type {
	case NUMBER:
		return p.ParseLiteral()
	case STRING:
		return p.ParseLiteral()
	case IDENTIFIER:
		return p.ParseIdentifier()
	default:
		panic("Expected a primary expression (number, string, or identifier)")
	}
}
