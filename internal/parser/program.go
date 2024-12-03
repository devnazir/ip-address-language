package parser

import (
	"os"
	"path/filepath"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseProgram() *ast.Program {
	filename := os.Args[1]
	rootFullDir := filepath.Dir(filename)

	program := &ast.Program{
		BaseNode: ast.BaseNode{
			Type:  ast.ProgramTree,
			Start: 0,
			End:   0,
		},
		Body:       []ast.ASTNode{},
		EntryPoint: rootFullDir,
	}

	lastPos := 0

	for p.pos < len(p.tokens) {

		if p.peek().Type == lx.TokenEOF {
			lastPos = p.peek().End
			break
		}

		p.ParseBodyProgram(program)
	}

	program.End = lastPos
	return program
}

func (p *Parser) ParseBodyProgram(program *ast.Program) (ast.ASTNode, error) {
	switch p.peek().Type {
	case lx.TokenKeyword:
		if p.peek().Value == lx.KeywordVar || p.peek().Value == lx.KeywordConst {
			declaration, err := p.ParseVariableDeclaration()

			if err != nil {
				panic(err)
			}
			program.Body = append(program.Body, declaration)
		}

		if p.peek().Value == lx.KeywordFunc {
			fnDeclaration, err := p.ParseFunctionDeclaration()
			if err != nil {
				panic(err)
			}
			program.Body = append(program.Body, fnDeclaration)
			p.next()
		}

	case lx.TokenShellKeyword:
		shellExpression, err := p.ParseShellExpression()
		if err != nil {
			panic(err)
		}
		program.Body = append(program.Body, shellExpression)

	case lx.TokenIdentifier:
		identifier, err := p.ParseTokenIdentifier()
		if err != nil {
			panic(err)
		}
		program.Body = append(program.Body, identifier)

	case lx.TokenSubshell:
		program.Body = append(program.Body, p.ParseSubShell())

	case lx.TokenEOF:
		return program, nil

	case lx.TokenSemicolon, lx.TokenComment, lx.TokenWhitespace:
		p.next()

	case lx.TokenIllegal:
		oops.IllegalTokenError(p.peek())

	default:
		return program, oops.SyntaxError(p.peek(), "Unknown token")
	}

	return program, nil
}

func (p *Parser) ParseTokenIdentifier() (ast.ASTNode, error) {
	identifier, err := p.ParseIdentifier()

	if err != nil {
		return nil, err
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

	default:
		return nil, oops.SyntaxError(p.peek(), "Unexpected token")
	}

	return nil, nil
}
