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
	KEYWORD        TokenType = "KEYWORD"
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
)

var keywords = Keywords{"if", "else", "func", "return", "var", "source"}

func generateKeywordPattern(keywords Keywords) string {
	return `\b(` + strings.Join(keywords, "|") + `)\b`
}

func compilePattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

var TokenSpecs = []TokenSpec{
	{KEYWORD, compilePattern(generateKeywordPattern(keywords))},
	{IDENTIFIER, compilePattern(`[a-zA-Z_]\w*`)},
	{NUMBER, compilePattern(`\b\d+(\.\d+)?\b`)},
	{OPERATOR, compilePattern(`[+\-*/=]`)},
	{LPAREN, compilePattern(`\(`)},
	{RPAREN, compilePattern(`\)`)},
	{LCURLY_BRACKET, compilePattern(`\{`)},
	{RCURLY_BRACKET, compilePattern(`\}`)},
	{SEMICOLON, compilePattern(`;`)},
	{STRING, compilePattern(`"[^"]*"`)},
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

type Lexer struct {
	Source string
	Tokens []Token
	Pos    int
}
