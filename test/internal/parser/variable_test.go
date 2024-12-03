package parser_test

import (
	"testing"

	"github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/internal/parser"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/stretchr/testify/assert"
)

func TestParseVariableDeclaration(t *testing.T) {
	t.Run("number variable declaration", func(t *testing.T) {
		text := "var x = 42;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   11,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 4,
					End:   11,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 4,
						End:   5,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 8,
						End:   10,
						Line:  1,
					},
					Value: 42,
					Raw:   "42",
				},
			},
			Kind:           "var",
			TypeAnnotation: "",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("string variable declaration", func(t *testing.T) {
		text := "var x = \"hello\";"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   16,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 4,
					End:   16,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 4,
						End:   5,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.StringLiteral{
					BaseNode: ast.BaseNode{
						Type:  "StringLiteral",
						Start: 8,
						End:   15,
						Line:  1,
					},
					Value: "hello",
					Raw:   "\"hello\"",
				},
			},
			Kind:           "var",
			TypeAnnotation: "",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("string template literal variable declaration", func(t *testing.T) {
		text := "var x = `hello $name`;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   22,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 4,
					End:   22,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 4,
						End:   5,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.StringTemplateLiteral{
					BaseNode: ast.BaseNode{
						Type:  "StringTemplateLiteral",
						Start: 8,
						End:   21,
						Line:  1,
					},
					Parts: []ast.ASTNode{
						ast.StringLiteral{
							BaseNode: ast.BaseNode{
								Type:  "StringLiteral",
								Start: 8,
								End:   13,
								Line:  1,
							},
							Value: "hello",
							Raw:   "hello ",
						},
						ast.Identifier{
							BaseNode: ast.BaseNode{
								Type:  "Identifier",
								Start: 14,
								End:   19,
								Line:  1,
							},
							Name:            "name",
							IsRestParameter: false,
						},
					},
				},
			},
			Kind:           "var",
			TypeAnnotation: "",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("single quote string variable declaration", func(t *testing.T) {
		text := "var x = 'hello';"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   16,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 4,
					End:   16,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 4,
						End:   5,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.StringLiteral{
					BaseNode: ast.BaseNode{
						Type:  "StringLiteral",
						Start: 8,
						End:   15,
						Line:  1,
					},
					Value: "'hello'",
					Raw:   "'hello'",
				},
			},

			Kind:           "var",
			TypeAnnotation: "",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("variable declaraion with subshell", func(t *testing.T) {
		text := "var x = $(ls);"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   14,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 4,
					End:   14,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 4,
						End:   5,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.SubShell{
					BaseNode: ast.BaseNode{
						Type:  "SubShell",
						Start: 8,
						End:   13,
						Line:  1,
					},
					Arguments: "ls",
				},
			},
			Kind:           "var",
			TypeAnnotation: "",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("variable declaraion without value", func(t *testing.T) {
		text := "var x;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		_, err := parser.NewParser(tokens).ParseVariableDeclaration()
		assert.NotNil(t, err)
	})

	t.Run("variable declaraion with type annotation", func(t *testing.T) {
		text := "var x string;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   13,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 4,
					End:   13,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 4,
						End:   5,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: nil,
			},
			Kind:           "var",
			TypeAnnotation: "string",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("variable declaraion with type annotation and value", func(t *testing.T) {
		text := "var x string = 42;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   18,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 4,
					End:   18,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 4,
						End:   5,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 15,
						End:   17,
						Line:  1,
					},
					Value: 42,
					Raw:   "42",
				},
			},
			Kind:           "var",
			TypeAnnotation: "string",
		}

		assert.Equal(t, expectedAst, actualAst)
	})
}

func TestConstantVariable(t *testing.T) {
	t.Run("constant variable declaration", func(t *testing.T) {
		text := "const x = 42;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   13,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 6,
					End:   13,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 6,
						End:   7,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 10,
						End:   12,
						Line:  1,
					},
					Value: 42,
					Raw:   "42",
				},
			},
			Kind:           "const",
			TypeAnnotation: "",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("constant variable without assignment", func(t *testing.T) {
		text := "const x;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		_, err := parser.NewParser(tokens).ParseVariableDeclaration()
		assert.NotNil(t, err)
	})

	t.Run("constant variable with type annotation and without assignment", func(t *testing.T) {
		text := "const x string;"
		tokens := lexer.NewLexer(text, "").Tokenize()
		_, err := parser.NewParser(tokens).ParseVariableDeclaration()
		assert.NotNil(t, err)
	})
}
