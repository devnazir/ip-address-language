package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseProgram() ast.Program {
	// mainDir := path.Dir(p.lexer.Filename)

	program := ast.Program{
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
		case KEYWORD:
			if p.peek().Value == lx.SOURCE {
				program.Body = append(program.Body, p.ParseSourceDeclaration())
			}

			if p.peek().Value == lx.VAR || p.peek().Value == lx.CONST {
				program.Body = append(program.Body, p.ParseVariableDeclaration())
			}

		case lx.SHELL_KEYWORD:
			program.Body = append(program.Body, p.ParseShellExpression())

		case IDENTIFIER:
			identToken := p.next()

			if p.peek().Type == OPERATOR && p.peek().Value == "=" {
				p.next()
				program.Body = append(program.Body, p.ParseAssignmentExpression(identToken))
			}

		case EOF:
			return program
		case lx.SEMICOLON, lx.COMMENT, lx.WHITESPACE:
			p.next()
		case ILLEGAL:
			oops.IllegalTokenError(p.peek())
			p.next()
		default:
			oops.UnexpectedTokenError(p.peek(), "")
			break
		}
	}

	return program
}
