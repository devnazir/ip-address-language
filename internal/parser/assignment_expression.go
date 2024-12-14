package parser

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
)

func (p *Parser) ParseAssignmentExpression(ident ast.Identifier) ast.ASTNode {
	expression := p.EvaluateAssignmentExpression()

	return ast.AssignmentExpression{
		Identifier: ast.Identifier{
			Name: ident.Name,
			BaseNode: ast.BaseNode{
				Type:  ast.AssignmentExpressionTree,
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

	// array expression
	if token.Type == lx.TokenLeftBracket {
		expr, err := p.ParseArrayExpression()
		if err != nil {
			panic(err)
		}
		return expr
	}

	// object expression
	if token.Type == lx.TokenLeftCurly {
		expr, err := p.ParseObjectExpression()
		if err != nil {
			panic(err)
		}
		return expr
	}

	for !endLoop && token.Type != lx.TokenShellKeyword {
		token := p.peek()

		if startLine != token.Line {
			// stop the loop
			endLoop = true
			break
		}

		switch token.Type {
		case
			lx.TokenNumber,
			lx.TokenIdentifier,
			lx.TokenDollarSign,
			lx.TokenTickQuote,
			lx.TokenDoubleQuote,
			lx.TokenBoolean:

			primaryExpression, err := p.ParsePrimaryExpression()
			if err != nil {
				panic(err)
			}

			output = append(output, primaryExpression)
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
				fnDeclaration, err := p.ParseFunctionDeclaration()
				if err != nil {
					panic(err)
				}
				output = append(output, fnDeclaration)
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

	if len(output) == 1 {
		return output[0]
	}

	return nil
}

func (p *Parser) ParsePrimaryExpression() (ast.ASTNode, error) {
	switch p.peek().Type {
	case lx.TokenNumber:
		return p.ParseNumberLiteral(), nil
	case lx.TokenDoubleQuote:
		return p.ParseStringLiteral(nil), nil
	case lx.TokenIdentifier, lx.TokenDollarSign:
		return p.ParseIdentifier(true)
	case lx.TokenTickQuote:
		return p.ParseStringTemplateLiteral()
	case lx.TokenBoolean:
		return p.ParseBooleanLiteral(), nil
	case lx.TokenLeftBracket:
		return p.ParseArrayExpression()
	case lx.TokenLeftCurly:
		return p.ParseObjectExpression()
	default:
		return nil, oops.SyntaxError(p.peek(), "Unexpected token"+p.peek().Value)
	}
}
