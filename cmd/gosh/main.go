package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/devnazir/gosh-script/internal/interpreter"
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/internal/parser"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func main() {

	if len(os.Args) < 2 {
		recovery := func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				debug.PrintStack()
			}
		}

		defer recovery()
		oops.ExpectedEntrypointFileError()
	}

	filename := os.Args[1]
	ext := filepath.Ext(filename)

	if ext != ".gsh" && ext != ".gosh" && ext != ".gs" {
		oops.InvalidFileExtensionError(filename)
	}

	lexer := lx.NewLexerFromFilename(filename)
	saveToFile("lexer.json", lexer) // for debug purpose

	tokens := lexer.Tokenize()
	saveToFile("tokens.json", lexer.Tokenize()) // for debug purpose

	parser := parser.NewParser(tokens, lexer)

	ast := parser.Parse()
	saveToFile("ast.json", ast) // for debug purpose

	interpreter := interpreter.NewInterpreter()
	interpreter.Interpret(ast)
}

func saveToFile(path string, data interface{}) {
	os.Mkdir("output", os.ModePerm)
	dir, _ := os.Getwd()
	fullPath := fmt.Sprintf("%s/output/%s", dir, path)

	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling data to JSON:", err)
		return
	}

	if _, err := file.Write(jsonData); err != nil {
		fmt.Println("Error writing JSON to file:", err)
	} else {
		// fmt.Printf("Data successfully saved to %s\n", fullPath)
	}
}
