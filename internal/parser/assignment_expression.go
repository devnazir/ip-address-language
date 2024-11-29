package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

func (p *Parser) ParseAssignmentExpression(ident *ast.Identifier) ast.ASTNode {
	expression := p.EvaluateAssignmentExpression()

	return ast.AssignmentExpression{
		Identifier: ast.Identifier{
			Name: ident.Name,
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.AssignmentExpression{}).Name(),
				Start: ident.Start,
				End:   ident.End,
				Line:  ident.Line,
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

	for !endLoop && token.Type != lx.TokenShellKeyword {
		token := p.peek()

		if startLine != token.Line {
			// stop the loop
			endLoop = true
			break
		}

		switch token.Type {
		case lx.TokenNumber, lx.TokenString, lx.TokenIdentifier, lx.TokenDollarSign, lx.TokenStringTemplateLiteral:
			output = append(output, p.ParsePrimaryExpression())
		case lx.TokenOperator:
			for len(operators) > 0 && Precedence[operators[len(operators)-1].Value] >= Precedence[token.Value] {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}

			operators = append(operators, token)
			p.next()
		case lx.TokenLeftParen:
			operators = append(operators, token)
			p.next()
		case lx.TokenRightParen:
			for operators[len(operators)-1].Type != lx.TokenLeftParen {
				output = append(output, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
				continue
			}

			operators = operators[:len(operators)-1]
			p.next()

		case lx.TokenSubshell:
			output = append(output, p.ParseSubShell())

		case lx.TokenKeyword:
			if token.Value == lx.KeywordFunc {
				output = append(output, p.ParseFunctionDeclaration())
				p.next()
				continue
			}

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

	if isBinaryExpression {
		return p.ParseBinaryExpression(output)
	}

	return output[0]
}

func (p *Parser) ParsePrimaryExpression() ast.ASTNode {
	switch p.peek().Type {
	case lx.TokenNumber:
		return p.ParseNumberLiteral()
	case lx.TokenString:
		return p.ParseStringLiteral(nil)
	case lx.TokenIdentifier:
		return p.ParseIdentifier()
	case lx.TokenDollarSign:
		return p.ParseIdentifier()
	case lx.TokenStringTemplateLiteral:
		return p.ParseStringTemplateLiteral()
	default:
		panic("Expected a primary expression (number, string, or identifier)")
	}
}
