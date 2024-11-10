package lexer

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func NewLexer(source, filename string) *Lexer {
	return &Lexer{Source: source, Tokens: []Token{}, Pos: 0, Filename: filename}
}

func NewLexerFromFilename(filename string) *Lexer {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", filename, err)
		panic(err)
	}

	return NewLexer(string(content), filename)
}

var TokensWithWhiteSpace []string = []string{
	ECHO,
}

func isTokenWithWhiteSpace(tokenValue string) bool {
	for _, token := range TokensWithWhiteSpace {
		if tokenValue == token {
			return true
		}
	}

	return false
}

func (l *Lexer) matchToken(chunk string, token *Token) bool {
	for _, spec := range TokenSpecs {
		if match := spec.pattern.FindString(chunk); match != "" {
			// fmt.Println("chunk", chunk, "match", match, "type", spec.tokenType)
			// Ensure the match occurs at the beginning of the chunk
			if strings.HasPrefix(chunk, match) {
				matchedValue := strings.TrimSpace(match)
				line := strings.Count(l.Source[:l.Pos], "\n") + 1

				token = &Token{
					Line:     line,
					Start:    l.Pos,
					End:      l.Pos + len(match),
					Value:    matchedValue,
					RawValue: match,
					Type:     spec.tokenType,
				}

				l.Tokens = append(l.Tokens, *token)
				l.Pos += len(match)

				return true
			}
		}
	}

	return false
}

var whitespaceRegex = regexp.MustCompile(`\s`)

func (l *Lexer) skipWhitespace() {
	for l.Pos < len(l.Source) && whitespaceRegex.MatchString(string(l.Source[l.Pos])) {
		l.Pos++
	}
}

func (l *Lexer) Tokenize() []Token {
	for l.Pos < len(l.Source) {
		l.skipWhitespace()

		token := &Token{}
		token.Start = l.Pos
		foundMatch := false

		// Try to match a valid token
		if l.matchToken(l.Source[l.Pos:], token) {
			foundMatch = true
		}

		// If no valid token is found, treat the characters as an illegal token
		if !foundMatch {
			illegalToken := ""

			// Capture all consecutive illegal characters (non-whitespace, non-token characters)
			for l.Pos < len(l.Source) && !strings.ContainsAny(string(l.Source[l.Pos]), " \t\n\r") {
				illegalToken += string(l.Source[l.Pos])
				l.Pos++
			}

			// Append the illegal token
			if strings.TrimSpace(illegalToken) != "" {
				line := strings.Count(l.Source[:l.Pos], "\n") + 1
				l.Tokens = append(l.Tokens, Token{Type: ILLEGAL, Start: token.Start, End: l.Pos, Value: illegalToken, Line: line})
			}
		}
	}

	if len(l.Tokens) == 0 || l.Tokens[len(l.Tokens)-1].Type != EOF {
		l.Tokens = append(l.Tokens, Token{Type: EOF, Start: l.Pos, End: l.Pos})
	}

	return l.Tokens
}
