package main

import (
	"encoding/json"
	"fmt"

	"log"

	"github.com/devnazir/gosh-script/pkg/lexer"
	"github.com/devnazir/gosh-script/pkg/parser"
)

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {
	lexer := lexer.NewLexerFromFilename(`./examples/variable.gsh`)
	tokens := lexer.Tokenize()
	parser := parser.NewParser(tokens, *lexer)

	jsonData, err := json.MarshalIndent(parser.Parse(), "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
