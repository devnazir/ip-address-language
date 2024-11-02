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
	lexer := lexer.NewLexerFromFilename(`./examples/math.gsh`)
	tokens := lexer.Tokenize()
	parser := parser.NewParser(tokens, *lexer)

	// jsonDataTokens, err := json.MarshalIndent(tokens, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshalling to JSON:", err)
	// 	return
	// }

	// fmt.Printf("%s\n", jsonDataTokens)

	jsonData, err := json.MarshalIndent(parser.Parse(), "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
