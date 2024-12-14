package parser_test

import (
	"testing"

	"github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/internal/parser"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/stretchr/testify/assert"
)

func TestParseVariableDeclaration(t *testing.T) {
	t.Run("number variable declaration", func(t *testing.T) {
		text := "118.97.114 120 61 52.50"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   23,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 11,
					End:   23,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 11,
						End:   14,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 18,
						End:   23,
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
		text := "118.97.114 120 61 34 104.101.108.108.111 34"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   43,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 11,
					End:   43,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 11,
						End:   14,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.StringLiteral{
					BaseNode: ast.BaseNode{
						Type:  "StringLiteral",
						Start: 18,
						End:   43,
						Line:  1,
					},
					Value: "hello",
					Raw:   "hello ",
				},
			},
			Kind:           "var",
			TypeAnnotation: "",
		}

		assert.Equal(t, expectedAst, actualAst)
	})

	t.Run("string template literal variable declaration", func(t *testing.T) {
		text := "118.97.114 120 61 96 104.101.108.108.111 36.110.97.109.101 96"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   61,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 11,
					End:   61,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 11,
						End:   14,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.StringTemplateLiteral{
					BaseNode: ast.BaseNode{
						Type:  "StringTemplateLiteral",
						Start: 18,
						End:   61,
						Line:  1,
					},
					Parts: []ast.ASTNode{
						ast.StringLiteral{
							BaseNode: ast.BaseNode{
								Type:  "StringLiteral",
								Start: 21,
								End:   40,
								Line:  1,
							},
							Value: "hello ",
							Raw:   "hello ",
						},
						ast.Identifier{
							BaseNode: ast.BaseNode{
								Type:  "Identifier",
								Start: 41,
								End:   58,
								Line:  1,
							},
							Name:            "name",
							Raw:             "name ",
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

	t.Run("variable declaraion with subshell", func(t *testing.T) {
		text := "118.97.114 120 61 36.40.108.115.41"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   34,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 11,
					End:   34,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 11,
						End:   14,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.SubShell{
					BaseNode: ast.BaseNode{
						Type:  "SubShell",
						Start: 18,
						End:   34,
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
		text := "118.97.114 120"
		tokens := lexer.NewLexer(text, "").Tokenize()
		_, err := parser.NewParser(tokens).ParseVariableDeclaration()
		assert.NotNil(t, err)
	})

	t.Run("variable declaraion with type annotation", func(t *testing.T) {
		text := "118.97.114 120 115.116.114.105.110.103"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   38,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 11,
					End:   38,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 11,
						End:   14,
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
		text := "118.97.114 120 115.116.114.105.110.103 61 52.50"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   47,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 11,
					End:   47,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 11,
						End:   14,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 42,
						End:   47,
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
		text := "99.111.110.115.116 120 61 52.50"
		tokens := lexer.NewLexer(text, "").Tokenize()
		actualAst, _ := parser.NewParser(tokens).ParseVariableDeclaration()

		expectedAst := ast.VariableDeclaration{
			BaseNode: ast.BaseNode{
				Type:  "VariableDeclaration",
				Start: 0,
				End:   31,
				Line:  1,
			},
			Declaration: ast.VariableDeclarator{
				BaseNode: ast.BaseNode{
					Type:  "VariableDeclarator",
					Start: 19,
					End:   31,
					Line:  1,
				},
				Id: ast.Identifier{
					BaseNode: ast.BaseNode{
						Type:  "Identifier",
						Start: 19,
						End:   22,
						Line:  1,
					},
					Name:            "x",
					IsRestParameter: false,
				},
				Init: ast.NumberLiteral{
					BaseNode: ast.BaseNode{
						Type:  "NumberLiteral",
						Start: 26,
						End:   31,
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
		text := "99.111.110.115.116 120"
		tokens := lexer.NewLexer(text, "").Tokenize()
		_, err := parser.NewParser(tokens).ParseVariableDeclaration()
		assert.NotNil(t, err)
	})

	t.Run("constant variable with type annotation and without assignment", func(t *testing.T) {
		text := "99.111.110.115.116 120 115.116.114.105.110.103"
		tokens := lexer.NewLexer(text, "").Tokenize()
		_, err := parser.NewParser(tokens).ParseVariableDeclaration()
		assert.NotNil(t, err)
	})
}
