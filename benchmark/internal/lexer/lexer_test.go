package lexer_test

import (
	"strings"
	"testing"

	"github.com/devnazir/ip-address-language/internal/lexer"
)

var heavyTestData = `
` + strings.Repeat("var x = 42;\n", 10000) + `

func foo() {
` + strings.Repeat("bar();\n", 10000) + `
}

func bar() {
` + strings.Repeat("echo $x;\n", 10000) + `
}

var longCalc = (1 + 2 - 3 * 4 / 5 % 6 + 7 - 8) * (9 + 10 - 11 * 12 / 13) % (14 + 15 - 16) * 17;
` + strings.Repeat("var z = longCalc + 42;\n", 10000) + `
`

func BenchmarkNewLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lx := lexer.NewLexer(heavyTestData, "")
		lx.Tokenize()
	}
}
