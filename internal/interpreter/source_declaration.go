package interpreter

import (
	"fmt"
	"strings"

	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/internal/parser"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/semantics"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (i *Interpreter) InterpretSourceDeclaration(sourceDeclaration ast.SourceDeclaration, entrypoint string) {
	for _, sources := range sourceDeclaration.Sources {
		file := sources.Value
		fileDir, err := utils.FindDirByFilename(entrypoint, file)
		alias := sources.Alias

		if err != nil {
			fmt.Println(err)
			return
		}

		lexer := lx.NewLexerFromFilename(fileDir + "/" + file)
		tokens := lexer.Tokenize()

		parser := parser.NewParser(tokens)
		ast := parser.Parse()
		i.InterpretSourceAst(ast, alias)
	}
}

func (i *Interpreter) InterpretSourceAst(p *ast.Program, alias string) {

	if len(alias) > 0 && !utils.IsAlpha(alias[0]) {
		oops.SourceAliasMustBeAlphanumericError(alias)
	}

	for _, nodeItem := range p.Body {
		switch nodeItem.GetType() {
		case ast.VariableDeclarationTree:
			node := nodeItem.(ast.VariableDeclaration)
			name := node.Declaration.Id.(ast.Identifier).Name

			value, _ := i.EvaluateVariableInit(node)

			if isIdentifierExported(name) {
				if len(alias) > 0 {
					i.symbolTable.Insert(alias, semantics.SymbolInfo{
						Kind: lx.KeywordSource,
						Value: map[string]interface{}{
							name: value,
						},
					})
					continue
				}

				i.InterpretVariableDeclaration(node)
			} else {
				i.InterpretNode(nodeItem, p.EntryPoint)
			}

		case ast.FunctionDeclarationTree:
			name := nodeItem.(ast.FunctionDeclaration).Identifier.Name
			params := nodeItem.(ast.FunctionDeclaration).Parameters

			if name == "init" {
				if len(params) > 0 {
					oops.InitFunctionCannotHaveParametersError(nodeItem.(ast.FunctionDeclaration))
				}

				i.InterpretBodyFunction(nodeItem.(ast.FunctionDeclaration), nil)
			}

		default:
			break
		}
	}
}

func isIdentifierExported(name string) bool {
	return strings.ToUpper(string(name[0])) == string(name[0])
}
