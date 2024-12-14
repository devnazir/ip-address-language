package parser

import (
	"reflect"

	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
)

func getPosition(node interface{}) (start, end int) {
	switch v := node.(type) {
	case ast.NumberLiteral:
		return v.Start, v.End
	case ast.StringLiteral:
		return v.Start, v.End
	}
	return 0, 0
}

func (p *Parser) ParseBinaryExpression(output []ast.ASTNode) ast.ASTNode {
	stack := []ast.ASTNode{}

	for _, nodeItem := range output {
		nodeType := reflect.TypeOf(nodeItem)

		if nodeType != reflect.TypeOf(lx.Token{}) {
			stack = append(stack, nodeItem)
			continue
		}

		right := stack[len(stack)-1]
		left := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		start, _ := getPosition(left)
		_, end := getPosition(right)

		stack = append(stack, ast.BinaryExpression{
			BaseNode: ast.BaseNode{
				Type:  ast.BinaryExpressionTree,
				Start: start,
				End:   end,
				Line:  nodeItem.(lx.Token).Line,
			},
			Operator: nodeItem.(lx.Token).Value,
			Left:     left,
			Right:    right,
		})
	}

	return stack[0]
}
