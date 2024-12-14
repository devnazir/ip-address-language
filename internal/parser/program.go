package parser

import (
	"os"
	"path/filepath"

	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
	"github.com/devnazir/ip-address-language/pkg/utils"
)

func (p *Parser) ParseProgram() *ast.Program {
	filename := os.Args[1]
	rootFullDir := filepath.Dir(filename)
	lastPos := 0

	program := &ast.Program{
		BaseNode: ast.BaseNode{
			Type:  ast.ProgramTree,
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		EntryPoint: rootFullDir,
	}

	body := ast.BodyProgram{}

	for p.pos < len(p.tokens) {
		if p.TokenTypeIs(lx.TokenEOF) {
			lastPos = p.peek().End
			break
		}

		result, err := p.ParseBodyProgram()
		if err != nil {
			panic(err)
		}

		body = append(body, result...)
	}

	program.Body = body
	program.End = lastPos
	return program
}

func (p *Parser) ParseBodyProgram() (ast.BodyProgram, error) {
	body := ast.BodyProgram{}

	switch p.peek().Type {
	case lx.TokenKeyword:
		keyword, err := p.ParseTokenKeyword()
		if err != nil {
			panic(err)
		}

		body = append(body, keyword)

	case lx.TokenShellKeyword:
		shellExpression, err := p.ParseShellExpression()
		if err != nil {
			panic(err)
		}
		body = append(body, shellExpression)

	case lx.TokenIdentifier:
		identifier, err := p.ParseTokenIdentifier()
		if err != nil {
			panic(err)
		}
		body = append(body, identifier)

	case lx.TokenSubshell:
		body = append(body, p.ParseSubShell())

	case lx.TokenEOF:
		return body, nil

	case lx.TokenSemicolon, lx.TokenComment, lx.TokenWhitespace:
		p.next()

	case lx.TokenIllegal:
		oops.IllegalTokenError(p.peek())

	default:
		utils.PrintJson(p.peek())
		panic(oops.SyntaxError(p.peek(), "Unknown token"))
	}

	return body, nil
}

func (p *Parser) ParseTokenKeyword() (ast.ASTNode, error) {
	switch p.peek().Value {
	case lx.KeywordVar, lx.KeywordConst:
		return p.ParseVariableDeclaration()

	case lx.KeywordFunc:
		return p.ParseFunctionDeclaration()

	case lx.KeywordIf:
		return p.ParseIfStatement()

	case lx.KeywordReturn:
		return p.ParseReturnStatement(ParseReturnStatementParams{
			isNextReturn: false,
		})

	default:
		return nil, oops.SyntaxError(p.peek(), "Unexpected token")
	}
}

func (p *Parser) ParseTokenIdentifier() (ast.ASTNode, error) {
	identifier, err := p.ParseIdentifier()

	if err != nil {
		return nil, err
	}

	if identifier.GetType() == ast.MemberExpressionTree {
		return identifier, nil
	}

	switch p.peek().Type {
	case lx.TokenOperator:
		if p.peek().Value == lx.EqualsSign {
			p.next()
			assignmentExpression := p.ParseAssignmentExpression(identifier.(ast.Identifier))
			return assignmentExpression, nil
		}

	case lx.TokenLeftParen:
		return p.parseCallExpression(identifier.(ast.Identifier)), nil

	case lx.TokenLeftBracket:
		memberExpression, err := p.ParseMemberExpression(identifier.(ast.Identifier))
		if err != nil {
			panic(err)
		}
		return memberExpression, nil

	case lx.TokenColon:
		p.next()

		if p.TokenValueIs(lx.EqualsSign) {
			p.next()
			assignmentExpression := p.ParseAssignmentExpression(identifier.(ast.Identifier))

			if ok := assignmentExpression.(ast.AssignmentExpression).Expression; ok == nil {
				return nil, oops.SyntaxError(p.peek(), "Expected value")
			}

			return ast.VariableDeclaration{
				BaseNode: ast.BaseNode{
					Type:  ast.VariableDeclarationTree,
					Start: identifier.(ast.Identifier).GetStart(),
					End:   identifier.(ast.Identifier).GetEnd(),
					Line:  identifier.GetLine(),
				},
				Declaration: ast.VariableDeclarator{
					BaseNode: ast.BaseNode{
						Type:  ast.VariableDeclaratorTree,
						Start: identifier.(ast.Identifier).GetStart(),
						End:   identifier.(ast.Identifier).GetEnd(),
						Line:  identifier.GetLine(),
					},
					Id:   identifier.(ast.Identifier),
					Init: assignmentExpression.(ast.AssignmentExpression).Expression,
				},
			}, nil
		}

	default:
		utils.PrintJson(p.peek())
		return nil, oops.SyntaxError(p.peek(), "Unexpected token")
	}

	return nil, nil
}
