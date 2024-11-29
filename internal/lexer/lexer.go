package lexer

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/devnazir/gosh-script/pkg/utils"
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

func (l *Lexer) matchPattern(pattern string) (string, bool) {
	re := regexp.MustCompile(pattern)
	match := re.FindString(l.Source[l.Pos:])
	if match != "" {
		l.Pos += len(match)
		return match, true
	}
	return "", false
}

func (l *Lexer) skipWhitespace() {
	for l.Pos < len(l.Source) {
		switch l.Source[l.Pos] {
		case ' ', '\t':
			l.Pos++
		case '\n':
			l.Line++
			l.Pos++
		default:
			return
		}
	}
}

func (l *Lexer) tokenizeWord(word string) TokenType {
	if t, ok := Keywords[word]; ok {
		return t
	}

	_, err := strconv.ParseFloat(word, 64)

	if err == nil {
		return TokenNumber
	}

	return TokenIdentifier
}

func (l *Lexer) tokenizeComment(symbole string) TokenType {
	if t, ok := CommentSymbols[symbole]; ok {
		if symbole == SingleLineComment {
			comment, _ := l.matchPattern(`.*\n?`)
			l.updateLineCount(comment)
		}

		isMultilineComment := symbole == MultiLineCommentStart || symbole == MultilineDocComment

		if isMultilineComment {
			comment, _ := l.matchPattern(`.*[\s\S]*?\*/`)
			l.updateLineCount(comment)
		}

		return t
	}

	return ""
}

func (l *Lexer) getWord() string {
	start := l.Pos
	for l.Pos < len(l.Source) && utils.IsAlphaNumeric(l.Source[l.Pos]) {
		l.Pos++
	}
	return l.Source[start:l.Pos]
}

func (l *Lexer) getComment() string {
	start := l.Pos
	nextPos := l.Pos + 1

	if nextPos >= len(l.Source) {
		return ""
	}

	startWithSlash := l.Source[l.Pos] == '/'
	if !startWithSlash {
		return ""
	}

	// if next character is not a * or /, it's not a comment
	if nextPos >= len(l.Source) || (l.Source[nextPos] != '*' && l.Source[nextPos] != '/') {
		return ""
	}

	for l.Pos < len(l.Source) && utils.IsComment(l.Source[l.Pos]) {
		l.Pos++
	}
	return l.Source[start:l.Pos]
}

func (l *Lexer) Tokenize() *[]Token {
	for l.Pos < len(l.Source) {
		l.skipWhitespace()

		startPos := l.Pos
		token := Token{Start: startPos, Line: l.Line}

		// Keyword or identifier
		word := l.getWord()

		if len(word) > 0 {
			token.Type = l.tokenizeWord(word)
			token.Value = strings.TrimSpace(word)
			token.End = l.Pos
			token.RawValue = l.Source[startPos:l.Pos]

			if l.Pos < len(l.Source) && l.Source[l.Pos] == ' ' {
				token.RawValue = l.Source[startPos : l.Pos+1]
			}

			l.Tokens = append(l.Tokens, token)
			continue
		}

		symbole := l.getComment()
		if symbole != "" {
			symbolType := l.tokenizeComment(symbole)
			if symbolType == TokenComment {
				continue
			}
		}

		// Match other patterns
		matched := false
		for typ, pattern := range TokenSpecs {
			if match, ok := l.matchPattern(pattern); ok {

				if typ == TokenComment || typ == TokenSemicolon {
					l.updateLineCount(match)
					matched = true
					continue
				}

				token.Type = typ
				token.Value = strings.TrimSpace(match)
				token.RawValue = l.Source[startPos:l.Pos]

				if l.Pos < len(l.Source) && l.Source[l.Pos] == ' ' {
					token.RawValue = l.Source[startPos : l.Pos+1]
				}

				token.End = l.Pos
				l.Tokens = append(l.Tokens, token)
				matched = true
				break
			}
		}

		// Illegal token handling
		if !matched && l.Pos < len(l.Source) {
			token.Type = TokenIllegal
			token.Value = string(l.Source[l.Pos])
			token.End = l.Pos + 1
			l.Tokens = append(l.Tokens, token)
			l.Pos++
		}
	}

	shouldAddEOF := len(l.Tokens) == 0 || l.Tokens[len(l.Tokens)-1].Type != TokenEOF

	if shouldAddEOF {
		l.Tokens = append(l.Tokens, Token{Type: TokenEOF, Start: l.Pos, End: l.Pos, Line: l.Line})
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
