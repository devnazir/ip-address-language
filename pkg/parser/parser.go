package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	lx "github.com/devnazir/gosh-script/pkg/lexer"
	"github.com/devnazir/gosh-script/pkg/node"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func NewParser(tokens []lx.Token, lexer *lx.Lexer) *Parser {
	return &Parser{
		tokens: tokens,
		lexer:  *lexer,
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

func (p *Parser) Parse() node.Program {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	program := p.ParseProgram()

	return program
}

func (p *Parser) ParseProgram() node.Program {
	// mainDir := path.Dir(p.lexer.Filename)

	program := node.Program{
		BaseNode: node.BaseNode{
			Type:  reflect.TypeOf(node.Program{}).Name(),
			Start: 0,
			End:   len(p.lexer.Source),
		},
		Body: []node.ASTNode{},
	}

	for p.pos < len(p.tokens) {

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

		case lx.SEMICOLON, lx.COMMENT:
			p.next()
		case ILLEGAL:
			oops.IllegalTokenError(p.peek())
		case IDENTIFIER:
			identToken := p.next()

			fmt.Println(identToken)

			if p.peek().Type == OPERATOR && p.peek().Value == "=" {
				p.next()
				program.Body = append(program.Body, p.ParseAssignmentExpression(identToken))
			} else {
				oops.UnexpectedTokenError(p.peek(), "")
			}
		case EOF:
			return program
		default:
			oops.UnexpectedTokenError(p.peek(), "")
		}
	}

	return program
}

func (p *Parser) ParseAssignmentExpression(identToken lx.Token) node.ASTNode {
	expression := p.EvaluateAssignmentExpression()

	return node.AssignmentExpression{
		Identifier: node.Identifier{
			Name: identToken.Value,
			BaseNode: node.BaseNode{
				Type:  reflect.TypeOf(node.AssignmentExpression{}).Name(),
				Start: identToken.Start,
				End:   identToken.End,
				Line:  identToken.Line,
			},
		},
		Expression: expression,
	}
}

func (p *Parser) ParseIdentifier() node.Identifier {
	trimmedName := strings.Trim(p.peek().Value, "$;")

	node := node.Identifier{
		Name: trimmedName,
		BaseNode: node.BaseNode{
			Type:  reflect.TypeOf(node.Identifier{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
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
func (p *Parser) ParseVariableDeclaration() node.VariableDeclaration {
	node := node.VariableDeclaration{
		BaseNode: node.BaseNode{
			Type:  "VariableDeclaration",
			Start: p.pos,
			End:   0,
			Line:  p.peek().Line,
		},
		Declarations: []node.VariableDeclarator{},
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

	node.Declarations = append(node.Declarations, p.ParseVariableDeclarator())

	p.next() // skip identifier, next to assignment operator or type annotation

	if p.peek().Type != PRIMITIVE_TYPE && varTypeToken.Value == lx.VAR {
		if p.peek().Type != OPERATOR && p.peek().Value != "=" {
			oops.ExpectedTypeAnnotationError(identToken)
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

		oops.UnexpectedTokenError(p.peek(), "=")
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
func (p *Parser) ParseLiteral() node.Literal {
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

	node := node.Literal{
		BaseNode: node.BaseNode{
			Type:  reflect.TypeOf(node.Literal{}).Name(),
			Start: p.peek().Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
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
func (p *Parser) ParseVariableDeclarator() node.VariableDeclarator {
	node := node.VariableDeclarator{
		BaseNode: node.BaseNode{
			Type:  reflect.TypeOf(node.VariableDeclarator{}).Name(),
			Start: p.peek().Start,
			End:   0,
			Line:  p.peek().Line,
		},
		Id: node.Identifier{
			Name: p.peek().Value,
			BaseNode: node.BaseNode{
				Type:  reflect.TypeOf(node.Identifier{}).Name(),
				Start: p.peek().Start,
				End:   p.peek().End,
				Line:  p.peek().Line,
			},
		},
		Init: nil,
	}

	return node
}

func (p *Parser) EvaluateAssignmentExpression() node.ASTNode {
	var output []node.ASTNode
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
func (p *Parser) ParseBinaryExpression(output []node.ASTNode) node.ASTNode {
	stack := []node.ASTNode{}

	for _, nodeItem := range output {
		nodeType := reflect.TypeOf(nodeItem)

		if nodeType != reflect.TypeOf(lx.Token{}) {
			stack = append(stack, nodeItem)
			continue
		}

		right := stack[len(stack)-1]
		left := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		stack = append(stack, node.BinaryExpression{
			BaseNode: node.BaseNode{
				Type:  reflect.TypeOf(node.BinaryExpression{}).Name(),
				Start: nodeItem.(lx.Token).Start,
				End:   nodeItem.(lx.Token).End,
				Line:  nodeItem.(lx.Token).Line,
			},
			Operator: nodeItem.(lx.Token).Value,
			Left:     left,
			Right:    right,
		})
	}

	return stack[0]
}

func (p *Parser) ParsePrimaryExpression() node.ASTNode {
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

func (p *Parser) ParseShellExpression() node.ASTNode {
	keyword := p.next()

	switch keyword.Value {
	case lx.ECHO:
		return p.ParseEchoStatement()
	default:
		oops.UnexpectedKeywordError(keyword)
	}

	return node.ShellExpression{}
}

func (p *Parser) ParseEchoStatement() node.ASTNode {
	startStmtToken := p.peek()
	arguments := []node.ASTNode{}
	flags := []string{}

	for p.peek().Type != lx.SEMICOLON && p.peek().Type != lx.EOF {
		switch p.peek().Type {
		case STRING, lx.IDENTIFIER:
			arguments = append(arguments, p.ParseLiteral())
		case lx.DOLLAR_SIGN:
			arguments = append(arguments, p.ParseIdentifier())
		case lx.FLAG:
			flags = append(flags, p.peek().Value)
			p.next()
		default:
			arguments = append(arguments, p.ParseLiteral())
		}
	}

	return node.ShellExpression{
		BaseNode: node.BaseNode{
			Type:  reflect.TypeOf(node.ShellExpression{}).Name(),
			Start: startStmtToken.Start,
			End:   startStmtToken.End,
			Line:  startStmtToken.Line,
		},
		Expression: node.EchoStatement{
			BaseNode: node.BaseNode{
				Type:  reflect.TypeOf(node.EchoStatement{}).Name(),
				Start: startStmtToken.Start,
				End:   startStmtToken.End,
				Line:  startStmtToken.Line,
			},
			Arguments: arguments,
			Flags:     flags,
		},
	}
}

func (p *Parser) ParseSourceDeclaration() node.ASTNode {
	token := p.next()
	sources := []node.ASTNode{}

	switch p.peek().Type {
	case STRING:
		sources = append(sources, p.ParseLiteral())

	case lx.LPAREN:
		p.next()
		endLoop := false

		for !endLoop {
			switch p.peek().Type {
			case STRING:
				sources = append(sources, p.ParseLiteral())
			case lx.RPAREN:
				endLoop = true
				p.next()
			default:
				oops.ExpectedTokenError(p.peek(), ")")
				p.next()
			}
		}
	default:
		oops.UnexpectedTokenError(p.peek(), "")
	}

	return node.SourceDeclaration{
		BaseNode: node.BaseNode{
			Type:  reflect.TypeOf(node.SourceDeclaration{}).Name(),
			Start: token.Start,
			End:   p.peek().End,
			Line:  p.peek().Line,
		},
		Sources: sources,
	}
}

func ParseToJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
