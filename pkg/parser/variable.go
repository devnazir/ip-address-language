package parser

import (
	"reflect"
	"strconv"
	"strings"

	lx "github.com/devnazir/gosh-script/pkg/lexer"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (p *Parser) ParseVariableDeclaration() VariableDeclaration {
	node := VariableDeclaration{
		BaseNode: BaseNode{
			Type:  "VariableDeclaration",
			Start: p.pos,
			End:   0,
		},
		Declarations: []VariableDeclarator{},
		Kind:         p.peek().Value,
	}

	p.next() // skip "var"
	// expect identifier
	if p.peek().Type != IDENTIFIER {

		if p.peek().Value != lx.VAR && p.peek().Value != lx.CONST {
			oops.IllegalIdentifier(p.peek())
		}

		oops.ExpectedIdentifier(p.peek())
	}

	node.Declarations = append(node.Declarations, p.ParseVariableDeclarator())
	p.next() // skip identifier, next to assignment operator

	// check if the next token has primitive type
	if p.peek().Type == PRIMITIVE_TYPE {
		primitiveType := p.peek().Value
		node.TypeAnnotation = primitiveType
		p.next()
	}

	// expect assignment operator
	operator := p.peek().Value
	if operator != "=" {
		oops.UnexpectedToken(p.peek(), "=")
	}

	p.next() // next to assignment expression
	node.Declarations[0].Init = p.ParseAssignmentExpression()
	node.Declarations[0].End = p.peek().End
	node.BaseNode.End = p.peek().End

	node.TypeAnnotation = p.ParseTypeAnnotation(node)

	return node
}

func (p *Parser) ParseLiteral() Literal {
	value := p.peek().Value
	var literalValue interface{}

	switch p.peek().Type {
	case NUMBER:
		if strings.Contains(value, ".") {
			literalValue, _ = strconv.ParseFloat(value, 64)
		} else {
			literalValue, _ = strconv.Atoi(value)
		}
	default:
		literalValue = value
	}

	node := Literal{
		BaseNode: BaseNode{
			Type:  "Literal",
			Start: p.peek().Start,
			End:   p.peek().End,
		},
		Value: literalValue,
		Raw:   value,
	}
	p.next()
	return node
}

func (p *Parser) ParseVariableDeclarator() VariableDeclarator {
	node := VariableDeclarator{
		BaseNode: BaseNode{
			Type:  "VariableDeclarator",
			Start: p.peek().Start,
			End:   0,
		},
		Id: Identifier{
			Name: p.peek().Value,
			BaseNode: BaseNode{
				Type:  "Identifier",
				Start: p.peek().Start,
				End:   p.peek().End,
			},
		},
		Init: nil,
	}

	return node
}

func (p *Parser) ParseAssignmentExpression() ASTNode {
	var output []ASTNode
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

func (p *Parser) ParseBinaryExpression(output []ASTNode) ASTNode {
	stack := []ASTNode{}

	for _, node := range output {
		nodeType := reflect.TypeOf(node)

		if nodeType != reflect.TypeOf(lx.Token{}) {
			stack = append(stack, node)
			continue
		}

		right := stack[len(stack)-1]
		left := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		stack = append(stack, BinaryExpression{
			BaseNode: BaseNode{
				Type:  "BinaryExpression",
				Start: node.(lx.Token).Start,
				End:   node.(lx.Token).End,
			},
			Operator: node.(lx.Token).Value,
			Left:     left,
			Right:    right,
		})
	}

	return stack[0]
}

func (p *Parser) ParsePrimaryExpression() ASTNode {
	switch p.peek().Type {
	case NUMBER:
		return p.ParseLiteral()
	case STRING:
		return p.ParseLiteral()
	case IDENTIFIER:
		// TODO: Implement identifier parsing
		return p.ParseLiteral()
	default:
		panic("Expected a primary expression (number, string, or identifier)")
	}
}
