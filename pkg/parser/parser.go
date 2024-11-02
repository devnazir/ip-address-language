package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	lx "github.com/devnazir/gosh-script/pkg/lexer"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func NewParser(tokens []lx.Token, lexer lx.Lexer) *Parser {
	return &Parser{
		tokens: tokens,
		lexer:  lexer,
		pos:    0,
	}
}

func (p *Parser) peek() lx.Token {
	if p.pos >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	}

	return p.tokens[p.pos]
}

func (p *Parser) next() lx.Token {
	if p.pos >= len(p.tokens) {
		return lx.Token{Type: lx.EOF}
	}

	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) Parse() ASTNode {
	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	return p.ParseProgram()
}

func (p *Parser) ParseProgram() Program {
	program := Program{
		BaseNode: BaseNode{
			Type:  reflect.TypeOf(Program{}).Name(),
			Start: 0,
			End:   len(p.lexer.Source),
		},
		Body: []ASTNode{},
	}

	for p.pos < len(p.tokens) {
		switch p.peek().Type {
		case KEYWORD:
			if p.peek().Value == lx.VAR || p.peek().Value == lx.CONST {
				program.Body = append(program.Body, p.ParseVariableDeclaration())
			} else {
				oops.UnexpectedKeyword(p.peek())
			}
		case lx.SEMICOLON, lx.COMMENT:
			p.next()
		case ILLEGAL:
			oops.IllegalToken(p.peek())
		case IDENTIFIER:
			identToken := p.next()

			if p.peek().Type == OPERATOR && p.peek().Value == "=" {
				p.next()
				program.Body = append(program.Body, p.ParseAssignmentExpression(identToken))
			}
		case EOF:
			return program
		default:
			oops.UnexpectedToken(p.peek(), "")
		}
	}

	return program
}

func (p *Parser) ParseAssignmentExpression(identToken lx.Token) ASTNode {
	expression := p.EvaluateAssignmentExpression()

	return AssignmentExpression{
		Identifier: Identifier{
			Name: identToken.Value,
			BaseNode: BaseNode{
				Type:  reflect.TypeOf(AssignmentExpression{}).Name(),
				Start: identToken.Start,
				End:   identToken.End,
			},
		},
		Expression: expression,
	}
}

func (p *Parser) ParseIdentifier() Identifier {
	node := Identifier{
		Name: p.peek().Value,
		BaseNode: BaseNode{
			Type:  reflect.TypeOf(Identifier{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
		},
	}
	p.next()
	return node
}

/*
ParseVariableDeclaration parses a variable declaration statement
and returns a VariableDeclaration node.
Example:

	var x = 10;
	var y = 20;
*/
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

	varTypeToken := p.next() // skip "var"
	identToken := p.peek()

	// expect identifier
	if p.peek().Type != IDENTIFIER {
		if p.peek().Value != lx.VAR && p.peek().Value != lx.CONST {
			oops.IllegalIdentifier(p.peek())
		}

		oops.ExpectedIdentifier(p.peek())
	}

	node.Declarations = append(node.Declarations, p.ParseVariableDeclarator())

	p.next() // skip identifier, next to assignment operator or type annotation

	if p.peek().Type != PRIMITIVE_TYPE && varTypeToken.Value == lx.VAR {
		if p.peek().Type != OPERATOR && p.peek().Value != "=" {
			oops.ExpectedTypeAnnotation(identToken)
		}
	}

	// check if the next token has primitive type
	if p.peek().Type == PRIMITIVE_TYPE {
		// primitiveType := p.peek().Value
		// node.TypeAnnotation = primitiveType
		p.next()
	}

	// expect assignment operator
	operator := p.peek().Value

	if p.peek().Type != OPERATOR && operator != "=" {

		// var can be used to declare a variable without assignment
		if varTypeToken.Value == lx.VAR {
			return node
		}

		oops.UnexpectedToken(p.peek(), "=")
	}

	p.next() // next to assignment expression
	node.Declarations[0].Init = p.EvaluateAssignmentExpression()
	node.Declarations[0].End = p.peek().End
	node.BaseNode.End = p.peek().End

	// node.TypeAnnotation = p.ParseTypeAnnotation(node)

	return node
}

/*
ParseLiteral parses a literal value and returns a Literal node.
Example:

	2 -> Literal{Value: 2}
	"hello" -> Literal{Value: "hello"}
*/
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
			Type:  reflect.TypeOf(Literal{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
		},
		Value: literalValue,
		Raw:   value,
	}
	p.next()
	return node
}

/*
ParseVariableDeclarator parses a variable declarator and returns a VariableDeclarator node.
*/
func (p *Parser) ParseVariableDeclarator() VariableDeclarator {
	node := VariableDeclarator{
		BaseNode: BaseNode{
			Type:  reflect.TypeOf(VariableDeclarator{}).Name(),
			Start: p.peek().Start,
			End:   0,
		},
		Id: Identifier{
			Name: p.peek().Value,
			BaseNode: BaseNode{
				Type:  reflect.TypeOf(Identifier{}).Name(),
				Start: p.peek().Start,
				End:   p.peek().End,
			},
		},
		Init: nil,
	}

	return node
}

func (p *Parser) EvaluateAssignmentExpression() ASTNode {
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
		case IDENTIFIER:
			output = append(output, p.ParsePrimaryExpression())
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

/*
ParseBinaryExpression parses a binary expression and returns a BinaryExpression node.
Example:

	2 + 3 -> BinaryExpression{Operator: "+", Left: Literal{Value: 2}, Right: Literal{Value: 3}}
	2 + 3 * 4 -> BinaryExpression{Operator: "+", Left: Literal{Value: 2}, Right: BinaryExpression{Operator: "*", Left: Literal{Value: 3}, Right: Literal{Value: 4}}}
*/
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
				Type:  reflect.TypeOf(BinaryExpression{}).Name(),
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
		return p.ParseIdentifier()
	default:
		panic("Expected a primary expression (number, string, or identifier)")
	}
}
