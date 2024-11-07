package main

import (
	"fmt"
	"os"

	"github.com/devnazir/gosh-script/pkg/interpreter"
	"github.com/devnazir/gosh-script/pkg/lexer"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/parser"
)

func main() {
	if len(os.Args) < 2 {
		recovery := func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}

		defer recovery()
		oops.ExpectedEntrypointFileError()
	}

	filename := os.Args[1]

	lexer := lexer.NewLexerFromFilename(filename)
	tokens := lexer.Tokenize()
	parser := parser.NewParser(tokens, lexer)

	ast := parser.Parse()
	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(ast)

	// jsonDataTokens, err := json.MarshalIndent(ast, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshalling to JSON:", err)
	// 	return
	// }

	// fmt.Printf("%s\n", jsonDataTokens)

}
