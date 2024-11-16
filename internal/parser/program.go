package parser

import (
	"fmt"
	"path/filepath"
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseProgram() *ast.Program {
	rootFullDir := filepath.Dir(p.lexer.Filename)

	program := &ast.Program{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.Program{}).Name(),
			Start: 0,
			End:   len(p.lexer.Source),
		},
		Body:       []ast.ASTNode{},
		EntryPoint: rootFullDir,
	}

	for p.pos < len(p.tokens) {

		if p.peek().Type == lx.TokenEOF {
			break
		}

		p.ParseBodyProgram(program)
	}

	return program
}

func (p *Parser) ParseBodyProgram(program *ast.Program) ast.ASTNode {
	switch p.peek().Type {
	case lx.TokenKeyword:
		if p.peek().Value == lx.KeywordVar || p.peek().Value == lx.KeywordConst {
			program.Body = append(program.Body, p.ParseVariableDeclaration())
		}

		if p.peek().Value == lx.KeywordFunc {
			program.Body = append(program.Body, p.ParseFunctionDeclaration())
			p.next()
		}

	case lx.TokenShellKeyword:
		program.Body = append(program.Body, p.ParseShellExpression())

	case lx.TokenIdentifier:
		program.Body = append(program.Body, p.ParseTokenIdentifier())

	case lx.TokenSubshell:
		program.Body = append(program.Body, p.ParseSubShell())

	case lx.TokenEOF:
		return program

	case lx.TokenSemicolon, lx.TokenComment, lx.TokenWhitespace:
		p.next()

	case lx.TokenIllegal:
		oops.IllegalTokenError(p.peek())

	default:
		oops.UnexpectedTokenError(p.peek(), "")
		return program
	}

	return program
}

func (p *Parser) ParseTokenIdentifier() ast.ASTNode {
	identifier := p.ParseIdentifier().(ast.Identifier)

	switch p.peek().Type {
	case lx.TokenOperator:
		if p.peek().Value == lx.EqualsSign {
			p.next()
			return p.ParseAssignmentExpression(&identifier)
		}

	case lx.TokenLeftParen:
		return p.parseCallExpression(&identifier)

	case lx.TokenLeftBracket:
		return p.ParseMemberExpression(&identifier)

	default:
		panic("Unexpected token " + fmt.Sprint(identifier.Name))
	}

	return nil
}
