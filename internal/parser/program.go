package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseProgram() *ast.Program {
	// mainDir := path.Dir(p.lexer.Filename)

	program := &ast.Program{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.Program{}).Name(),
			Start: 0,
			End:   len(p.lexer.Source),
		},
		Body: []ast.ASTNode{},
	}

	for p.pos < len(p.tokens) {
		// fmt.Println(p.peek())
		switch p.peek().Type {
		case lx.TokenKeyword:
			if p.peek().Value == lx.KeywordSource {
				program.Body = append(program.Body, p.ParseSourceDeclaration())
			}

			if p.peek().Value == lx.KeywordVar || p.peek().Value == lx.KeywordConst {
				program.Body = append(program.Body, p.ParseVariableDeclaration())
			}

		case lx.TokenShellKeyword:
			program.Body = append(program.Body, p.ParseShellExpression())

		case lx.TokenIdentifier:
			identToken := p.next()

			if p.peek().Type == lx.TokenOperator && p.peek().Value == "=" {
				p.next()
				program.Body = append(program.Body, p.ParseAssignmentExpression(&identToken))
			}

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
	}

	return program
}
