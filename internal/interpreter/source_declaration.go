package interpreter

import (
	"fmt"
	"strings"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/internal/parser"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (i *Interpreter) InterpretSourceDeclaration(sources []ast.Source, entrypoint string) {
	for _, sources := range sources {
		file := sources.Value
		fileDir, err := utils.FindDirByFilename(entrypoint, file)
		alias := sources.Alias

		if err != nil {
			fmt.Println(err)
			return
		}

		lexer := lx.NewLexerFromFilename(fileDir + "/" + file)
		tokens := lexer.Tokenize()

		parser := parser.NewParser(tokens, lexer)
		ast := parser.Parse()
		i.InterpretSourceAst(ast, alias)
	}
}

func (i *Interpreter) InterpretSourceAst(p *ast.Program, alias string) {
	for _, nodeItem := range p.Body {
		switch nodeItem.(type) {
		case *ast.VariableDeclaration:
			name := nodeItem.(*ast.VariableDeclaration).Declaration.Id.(ast.Identifier).Name

			if isIdentifierExported(name) {

				if alias != "" {
					name = alias + "." + name
				}

				node := nodeItem.(*ast.VariableDeclaration)

				i.InterpretVariableDeclaration(ast.VariableDeclaration{
					Declaration: ast.VariableDeclarator{
						Id: ast.Identifier{
							Name:     name,
							BaseNode: node.Declaration.Id.(ast.Identifier).BaseNode,
						},
						BaseNode: node.Declaration.BaseNode,
						Init:     node.Declaration.Init,
					},
					BaseNode:       node.BaseNode,
					Kind:           node.Kind,
					TypeAnnotation: node.TypeAnnotation,
				})
			} else {
				i.InterpretNode(nodeItem, p.EntryPoint)
			}

		case *ast.FunctionDeclaration:
			name := nodeItem.(*ast.FunctionDeclaration).Identifier.Name
			params := nodeItem.(*ast.FunctionDeclaration).Parameters

			if name == "init" {

				if len(params) > 0 {
					panic("init function cannot have parameters")
				}

				i.InterpretBodyFunction(nodeItem.(ast.FunctionDeclaration), nil)
			}

			// default:
			// 	InterpretNode(i, nodeItem, p.EntryPoint)
		}
	}
}

func isIdentifierExported(name string) bool {
	return strings.ToUpper(string(name[0])) == string(name[0])
}
