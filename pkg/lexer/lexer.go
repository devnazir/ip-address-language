package lexer

import (
	"os"
	"strings"
)

func NewLexer(source string) *Lexer {
	return &Lexer{Source: source, Tokens: []Token{}, Pos: 0}
}

func NewLexerFromFilename(filename string) *Lexer {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, _ := file.Stat()
	fileSize := stat.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}

	return NewLexer(string(buffer))
}

func (l *Lexer) matchToken(chunk string, token *Token) bool {
	for _, spec := range TokenSpecs {
		if match := spec.pattern.FindString(chunk); match != "" {
			// Ensure the match occurs at the beginning of the chunk
			if strings.HasPrefix(chunk, match) {
				line := strings.Count(l.Source[:l.Pos], "\n") + 1
				token.Line = line
				token.Start = l.Pos
				token.End = l.Pos + len(match)
				token.Value = strings.TrimSpace(match)
				token.Type = spec.tokenType
				l.Tokens = append(l.Tokens, *token)
				l.Pos += len(match)
				return true
			}
		}
	}

	return false
}

func (l *Lexer) skipWhitespace() {
	for l.Pos < len(l.Source) {
		char := l.Source[l.Pos]
		// Check for whitespace and newline characters
		if char == ' ' || char == '\n' || char == '\t' || char == '\r' {
			l.Pos++
		} else {
			break
		}
	}
}

func (l *Lexer) Tokenize() []Token {
	for l.Pos < len(l.Source) {
		// Skip any whitespace characters first
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

	// Append EOF token at the end
	l.Tokens = append(l.Tokens, Token{Type: EOF, Start: l.Pos, End: l.Pos})
	return l.Tokens
}
