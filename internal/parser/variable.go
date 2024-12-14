package parser

import (
	"fmt"

	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
)

func (p *Parser) ParseVariableDeclaration() (ast.VariableDeclaration, error) {
	tree := ast.VariableDeclaration{
		BaseNode: ast.BaseNode{
			Type:  ast.VariableDeclarationTree,
			Start: p.pos,
			End:   0,
			Line:  p.peek().Line,
		},
		Declaration: ast.VariableDeclarator{},
		Kind:        p.peek().Value,
	}

	varTypeToken := p.next() // skip "var"
	identToken := p.peek()

	// expect identifier
	if p.peek().Type != lx.TokenIdentifier {
		if p.peek().Value != lx.KeywordVar && p.peek().Value != lx.KeywordConst {
			return ast.VariableDeclaration{}, fmt.Errorf(oops.CreateErrorMessage(p.peek(), "Illegal identifier"))
		}

		return ast.VariableDeclaration{}, fmt.Errorf(oops.CreateErrorMessage(p.peek(), "Expected identifier"))
	}

	tree.Declaration = p.ParseVariableDeclarator()

	p.next() // skip identifier, next to assignment operator or type annotation

	if p.peek().Type != lx.TokenPrimitiveType && varTypeToken.Value == lx.KeywordVar {
		if p.peek().Type != lx.TokenOperator && p.peek().Value != "=" {
			return ast.VariableDeclaration{}, fmt.Errorf(oops.CreateErrorMessage(identToken, "Expected assignment operator"))
		}
	}

	// check if the next token has primitive type
	if p.peek().Type == lx.TokenPrimitiveType {
		primitiveType := p.peek().Value
		tree.TypeAnnotation = primitiveType
		p.next()
	}

	// expect assignment operator
	operator := p.peek().Value

	if p.peek().Type != lx.TokenOperator && operator != "=" {

		// var can be used to declare a variable without assignment
		if varTypeToken.Value == lx.KeywordVar {
			tree.Declaration.End = p.peek().End
			tree.BaseNode.End = p.peek().End
			return tree, nil
		}

		return ast.VariableDeclaration{}, fmt.Errorf(oops.CreateErrorMessage(p.peek(), "Expected assignment operator"))
	}

	p.next() // next to assignment expression
	tree.Declaration.Init = p.EvaluateAssignmentExpression()
	tree.Declaration.End = p.peek().End
	tree.BaseNode.End = p.peek().End

	return tree, nil
}

func (p *Parser) ParseVariableDeclarator() ast.VariableDeclarator {
	tree := ast.VariableDeclarator{
		BaseNode: ast.BaseNode{
			Type:  ast.VariableDeclaratorTree,
			Start: p.peek().Start,
			End:   0,
			Line:  p.peek().Line,
		},
		Id: ast.Identifier{
			Name: p.peek().Value,
			BaseNode: ast.BaseNode{
				Type:  ast.IdentifierTree,
				Start: p.peek().Start,
				End:   p.peek().End,
				Line:  p.peek().Line,
			},
		},
		Init: nil,
	}

	return tree
}
