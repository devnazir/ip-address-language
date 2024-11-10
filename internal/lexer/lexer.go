package lexer

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func NewLexer(source, filename string) *Lexer {
	return &Lexer{Source: source, Tokens: []Token{}, Pos: 0, Filename: filename, Line: 1}
}

func NewLexerFromFilename(filename string) *Lexer {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s, %v\n", filename, err)
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %s, %v\n", filename, err)
		panic(err)
	}

	return NewLexer(string(content), filename)
}

func (l *Lexer) matchToken(chunk string, token *Token) bool {
	for _, spec := range tokenSpecs {
		if match := spec.Pattern.FindString(chunk); match != "" && match == chunk[:len(match)] {
			matchedValue := strings.TrimSpace(match)

			token.Line = l.Line
			token.Start = l.Pos
			token.End = l.Pos + len(match)
			token.Value = matchedValue
			token.RawValue = match
			token.Type = spec.Type

			l.Tokens = append(l.Tokens, *token)
			l.Pos += len(match)

			l.updateLineCount(match)
			return true
		}
	}
	return false
}

var whitespaceRegex = regexp.MustCompile(`\s`)

func (l *Lexer) skipWhitespace() {
	for l.Pos < len(l.Source) && whitespaceRegex.MatchString(string(l.Source[l.Pos])) {
		if l.Source[l.Pos] == '\n' {
			l.Line++
		}
		l.Pos++
	}
}

func (l *Lexer) Tokenize() *[]Token {
	for l.Pos < len(l.Source) {
		l.skipWhitespace()

		token := &Token{}
		token.Start = l.Pos

		if l.matchToken(l.Source[l.Pos:], token) {
			continue
		}

		startPos := l.Pos
		for l.Pos < len(l.Source) && !whitespaceRegex.MatchString(string(l.Source[l.Pos])) {
			l.Pos++
		}

		if startPos < l.Pos {
			illegalToken := l.Source[startPos:l.Pos]
			l.Tokens = append(l.Tokens, Token{
				Type:     TokenIllegal,
				Start:    startPos,
				End:      l.Pos,
				Value:    illegalToken,
				RawValue: illegalToken,
				Line:     l.Line,
			})
		}
	}

	if len(l.Tokens) == 0 || l.Tokens[len(l.Tokens)-1].Type != TokenEOF {
		l.Tokens = append(l.Tokens, Token{Type: TokenEOF, Start: l.Pos, End: l.Pos})
	}

	return &l.Tokens
}

func (l *Lexer) updateLineCount(text string) {
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			l.Line++
		}
	}
}
