package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/devnazir/gosh-script/internal/interpreter"
	"github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/internal/parser"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func main() {
	text := "Hello World Gdas $hello1$hello2$hello3 $hello4 asdasdas"

	// Match all text that does not start with `$`
	nonVarRe := regexp.MustCompile(`(?:^|[^$])\b([A-Za-z\s]+)\b(?:[^$]|$)`)
	// Match all variables that start with `$`
	varRe := regexp.MustCompile(`\$\w+`)

	// Find all non-variable parts
	nonVarMatches := nonVarRe.FindAllString(text, -1)
	// Join non-variable parts into a single string
	nonVarText := ""
	for _, match := range nonVarMatches {
		nonVarText += match
	}

	// Find all variable parts
	varMatches := varRe.FindAllString(text, -1)
	// Join variable parts into a single string
	varText := ""
	for _, match := range varMatches {
		varText += match + " "
	}

	// Print the results
	fmt.Println("String is:", nonVarText)
	fmt.Println("Vars are:", varText)

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
