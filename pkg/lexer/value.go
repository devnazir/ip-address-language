package lexer

import (
	"regexp"
	"strings"
)

type TokenType string
type Keywords []string

type TokenSpec struct {
	tokenType TokenType
	pattern   *regexp.Regexp
}

const (
	IDENTIFIER     TokenType = "IDENTIFIER"
	PRIMITIVE_TYPE TokenType = "PRIMITIVE_TYPE"
	COMPOSITE_TYPE TokenType = "COMPOSITE_TYPE"
	KEYWORD        TokenType = "KEYWORD"
	SHELL_KEYWORD  TokenType = "SHELL_KEYWORD"
	NUMBER         TokenType = "NUMBER"
	OPERATOR       TokenType = "OPERATOR"
	LPAREN         TokenType = "LPAREN"
	RPAREN         TokenType = "RPAREN"
	LCURLY_BRACKET TokenType = "LCURLY_BRACKET"
	RCURLY_BRACKET TokenType = "RCURLY_BRACKET"
	SEMICOLON      TokenType = "SEMICOLON"
	EOF            TokenType = "EOF"
	STRING         TokenType = "STRING"
	FUNC           TokenType = "FUNC"
	RETURN         TokenType = "RETURN"
	ILLEGAL        TokenType = "ILLEGAL"
	COMMENT        TokenType = "COMMENT"
	DOLLAR_SIGN    TokenType = "DOLLAR_SIGN"
	FLAG           TokenType = "FLAG"
	COMMA          TokenType = "COMMA"
)

const (
	VAR    = "var"
	CONST  = "const"
	ECHO   = "echo"
	SOURCE = "source"
)

var variableKeywords = Keywords{
	VAR,
	CONST,
}

var keywords = append(variableKeywords, Keywords{
	"if",
	"else",
	"func",
	"return",
	"source",
}...)

var shellKeywords = Keywords{
	ECHO,
}

var primitiveTypes = Keywords{
	// Boolean types
	"bool",

	// Numeric types
	"int",
	"float64",

	// String type
	"string",
}

func generatePattern(keywords []string) string {
	return `(\s*|^)` + `\b(` + strings.Join(keywords, "|") + `)\b(\s*|$)`
}

func compilePattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

var TokenSpecs = []TokenSpec{
	{COMMENT, compilePattern(`(//.*)|(/\*[\s\S]*?\*/)`)},
	{KEYWORD, compilePattern(generatePattern(keywords))},
	{SHELL_KEYWORD, compilePattern(generatePattern(shellKeywords))},
	{PRIMITIVE_TYPE, compilePattern(generatePattern(primitiveTypes))},

	{IDENTIFIER, compilePattern(`\b[a-zA-Z_][a-zA-Z0-9_]*\b`)},
	{FLAG, compilePattern(`-\w+`)},
	{NUMBER, compilePattern(`\b\d+(\.\d+)?\b`)},
	{OPERATOR, compilePattern(`[+\-*/=]`)},
	{STRING, compilePattern(`"[^"]*"`)},

	{DOLLAR_SIGN, compilePattern(`\$\w+`)},
	{LPAREN, compilePattern(`\(`)},
	{RPAREN, compilePattern(`\)`)},
	{LCURLY_BRACKET, compilePattern(`\{`)},
	{RCURLY_BRACKET, compilePattern(`\}`)},
	{SEMICOLON, compilePattern(`;`)},
	{COMMA, compilePattern(`,`)},
}

func TokenMap() map[TokenType]*regexp.Regexp {
	tokenMap := make(map[TokenType]*regexp.Regexp)
	for _, spec := range TokenSpecs {
		tokenMap[spec.tokenType] = spec.pattern
	}
	return tokenMap
}

type Token struct {
	Type  TokenType
	Start int
	End   int
	Value string
	Line  int
}

func (t Token) GetLine() int {
	return t.Line
}

type Lexer struct {
	Source   string
	Tokens   []Token
	Pos      int
	Filename string
}
