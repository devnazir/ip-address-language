package main

import (
	"encoding/json"
	"fmt"

	"github.com/devnazir/gosh-script/pkg"
)

func main() {
	lexer := pkg.NewLexer(`var test = 10`)
	tokens := lexer.Tokenize()
	parser := pkg.NewParser(tokens, *lexer)

	jsonData, err := json.MarshalIndent(parser.Parse(), "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
