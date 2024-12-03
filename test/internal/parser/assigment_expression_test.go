package parser_test

import (
	"testing"

	"github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/internal/parser"
	"github.com/devnazir/gosh-script/pkg/ast"
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
			input: GenerateSingleAST("x = 42;"),
			expected: ast.AssignmentExpression{
				Identifier: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 0,
						End:   1,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Expression: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 4,
						End:   6,
						Line:  1,
					},
					Value: 42,
					Raw:   "42",
				},
			},
		},
		{
			name:  "binary expression",
			input: GenerateSingleAST("x = 42 + 42"),
			expected: ast.AssignmentExpression{
				Identifier: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 0,
						End:   1,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Expression: ast.BinaryExpression{
					BaseNode: ast.BaseNode{
						Type:  "BinaryExpression",
						Start: 4,
						End:   11,
						Line:  1,
					},
					Operator: "+",
					Left: ast.NumberLiteral{
						BaseNode: ast.BaseNode{
							Type:  "NumberLiteral",
							Start: 4,
							End:   6,
							Line:  1,
						},
						Value: 42,
						Raw:   "42",
					},
					Right: ast.NumberLiteral{
						BaseNode: ast.BaseNode{
							Type:  "NumberLiteral",
							Start: 9,
							End:   11,
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
			input: GenerateSingleAST("x = func() { }"),
			expected: ast.AssignmentExpression{
				Identifier: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 0,
						End:   1,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Expression: ast.FunctionDeclaration{
					BaseNode: ast.BaseNode{
						Type:  "FunctionDeclaration",
						Start: 4,
						End:   8,
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
