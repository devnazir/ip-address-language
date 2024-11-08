package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseVariableDeclaration() ast.VariableDeclaration {
	ast := ast.VariableDeclaration{
		BaseNode: ast.BaseNode{
			Type:  "VariableDeclaration",
			Start: p.pos,
			End:   0,
			Line:  p.peek().Line,
		},
		Declarations: []ast.VariableDeclarator{},
		Kind:         p.peek().Value,
	}

	varTypeToken := p.next() // skip "var"
	identToken := p.peek()

	// expect identifier
	if p.peek().Type != IDENTIFIER {
		if p.peek().Value != lx.VAR && p.peek().Value != lx.CONST {
			oops.IllegalIdentifierError(p.peek())
		}

		oops.ExpectedIdentifierError(p.peek())
	}

	ast.Declarations = append(ast.Declarations, p.ParseVariableDeclarator())

	p.next() // skip identifier, next to assignment operator or type annotation

	if p.peek().Type != PRIMITIVE_TYPE && varTypeToken.Value == lx.VAR {
		if p.peek().Type != OPERATOR && p.peek().Value != "=" {
			oops.ExpectedTypeAnnotationError(identToken)
		}
	}

	// check if the next token has primitive type
	if p.peek().Type == PRIMITIVE_TYPE {
		// primitiveType := p.peek().Value
		// ast.TypeAnnotation = primitiveType
		p.next()
	}

	// expect assignment operator
	operator := p.peek().Value

	if p.peek().Type != OPERATOR && operator != "=" {

		// var can be used to declare a variable without assignment
		if varTypeToken.Value == lx.VAR {
			return ast
		}

		oops.UnexpectedTokenError(p.peek(), "=")
	}

	p.next() // next to assignment expression
	ast.Declarations[0].Init = p.EvaluateAssignmentExpression()
	ast.Declarations[0].End = p.peek().End
	ast.BaseNode.End = p.peek().End

	// ast.TypeAnnotation = p.ParseTypeAnnotation(ast)

	return ast
}

func (p *Parser) ParseVariableDeclarator() ast.VariableDeclarator {
	ast := ast.VariableDeclarator{
		BaseNode: ast.BaseNode{
			Type:  reflect.TypeOf(ast.VariableDeclarator{}).Name(),
			Start: p.peek().Start,
			End:   0,
			Line:  p.peek().Line,
		},
		Id: ast.Identifier{
			Name: p.peek().Value,
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.Identifier{}).Name(),
				Start: p.peek().Start,
				End:   p.peek().End,
				Line:  p.peek().Line,
			},
		},
		Init: nil,
	}

	return ast
}