package parser

import (
	"reflect"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
)

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

		stack = append(stack, ast.BinaryExpression{
			BaseNode: ast.BaseNode{
				Type:  reflect.TypeOf(ast.BinaryExpression{}).Name(),
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
