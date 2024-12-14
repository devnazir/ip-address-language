package parser_test

import (
	"testing"

	"github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/internal/parser"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/stretchr/testify/assert"
)

func GenerateSingleAST(value string) ast.AssignmentExpression {
	tokens := lexer.NewLexer(value, "").Tokenize()
	result := parser.NewParser(tokens).Parse()
	return result.Body[0].(ast.AssignmentExpression)
}

func TestParseAssignmentExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    ast.AssignmentExpression
		expected ast.AssignmentExpression
	}{
		{
			name:  "simple assignment",
			input: GenerateSingleAST("120 61 52.50 59"),
			expected: ast.AssignmentExpression{
				Identifier: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "AssignmentExpression",
						Start: 0,
						End:   3,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Expression: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 7,
						End:   12,
						Line:  1,
					},
					Value: 42,
					Raw:   "42",
				},
			},
		},
		{
			name:  "binary expression",
			input: GenerateSingleAST("120 61 52.50 43 52.50"),
			expected: ast.AssignmentExpression{
				Identifier: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "AssignmentExpression",
						Start: 0,
						End:   3,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Expression: ast.BinaryExpression{
					BaseNode: ast.BaseNode{
						Type:  "BinaryExpression",
						Start: 7,
						End:   21,
						Line:  1,
					},
					Operator: "+",
					Left: ast.NumberLiteral{
						BaseNode: ast.BaseNode{
							Type:  "NumberLiteral",
							Start: 7,
							End:   12,
							Line:  1,
						},
						Value: 42,
						Raw:   "42",
					},
					Right: ast.NumberLiteral{
						BaseNode: ast.BaseNode{
							Type:  "NumberLiteral",
							Start: 16,
							End:   21,
							Line:  1,
						},
						Value: 42,
						Raw:   "42",
					},
				},
			},
		},
		{
			name:  "anonymous function",
			input: GenerateSingleAST("120 61 102.117.110.99 40 41 123 125"),
			expected: ast.AssignmentExpression{
				Identifier: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "AssignmentExpression",
						Start: 0,
						End:   3,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Expression: ast.FunctionDeclaration{
					BaseNode: ast.BaseNode{
						Type:  "FunctionDeclaration",
						Start: 7,
						End:   21,
						Line:  1,
					},
					Identifier:  ast.Identifier{},
					Parameters:  []ast.Identifier{},
					IsAnonymous: true,
					Body:        []ast.ASTNode{},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.input)
		})
	}
}
