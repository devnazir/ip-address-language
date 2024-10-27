package pkg

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

var keywords = Keywords{
	"if",
	"else",
	"return",
	"var",
	"func",
	"return",
	"source",
}

var tokenSpecs = []TokenSpec{
	{KEYWORD, compilePattern(`\b(if|else|func|return|var|source)\b`)},
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
	for _, spec := range tokenSpecs {
		tokenMap[spec.tokenType] = spec.pattern
	}
	return tokenMap
}

type Token struct {
	Type  TokenType
	Start int
	End   int
	Value string
}

type Lexer struct {
	source string
	tokens []Token
	pos    int
}

func compilePattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

// create lexer
func NewLexer(source string) *Lexer {
	return &Lexer{source: source, tokens: []Token{}, pos: 0}
}

func (l *Lexer) matchToken(chunk string, token *Token) bool {
	for _, spec := range tokenSpecs {
		if match := spec.pattern.FindString(chunk); match != "" {
			// Ensure the match occurs at the beginning of the chunk
			if strings.HasPrefix(chunk, match) {
				token.End = l.pos + len(match)
				token.Value = match
				token.Type = spec.tokenType
				l.tokens = append(l.tokens, *token)
				l.pos += len(match)
				return true
			}
		}
	}

	return false
}

func (l *Lexer) skipWhitespace(chunk string) {
	for _, char := range chunk {
		// check whitespace character
		if char == ' ' || char == '\n' || char == '\t' || char == '\r' {
			l.pos++
		} else {
			break
		}
	}
}

func (l *Lexer) Tokenize() []Token {
	for l.pos < len(l.source) {
		l.skipWhitespace(l.source[l.pos:])

		chunk := l.source[l.pos:]

		token := &Token{}
		token.Start = l.pos
		foundMatch := false

		if match := l.matchToken(chunk, token); match != false {
			foundMatch = true
		}

		if !foundMatch {
			var illegalTokenName string
			if l.pos < len(l.source) {
				illegalTokenName = strings.TrimSpace(string(l.source[l.pos]))
			} else {
				illegalTokenName = ""
			}
			l.tokens = append(l.tokens, Token{Type: ILLEGAL, Start: l.pos, End: l.pos + 1, Value: illegalTokenName})
			l.pos++
			break
		}
	}

	// Append EOF token at the end
	l.tokens = append(l.tokens, Token{Type: EOF, Start: l.pos, End: l.pos})
	return l.tokens
}
