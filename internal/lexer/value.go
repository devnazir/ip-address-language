package lexer

import (
	"regexp"
	"strings"
)

type TokenType string
type KeywordList []string

type TokenSpec struct {
	Type    TokenType
	Pattern *regexp.Regexp
}

const (
	TokenIdentifier    TokenType = "IDENTIFIER"
	TokenPrimitiveType TokenType = "PRIMITIVE_TYPE"
	TokenCompositeType TokenType = "COMPOSITE_TYPE"
	TokenKeyword       TokenType = "KEYWORD"
	TokenShellKeyword  TokenType = "SHELL_KEYWORD"
	TokenNumber        TokenType = "NUMBER"
	TokenOperator      TokenType = "OPERATOR"
	TokenLeftParen     TokenType = "LEFT_PAREN"
	TokenRightParen    TokenType = "RIGHT_PAREN"
	TokenLeftCurly     TokenType = "LEFT_CURLY_BRACKET"
	TokenRightCurly    TokenType = "RIGHT_CURLY_BRACKET"
	TokenSemicolon     TokenType = "SEMICOLON"
	TokenColon         TokenType = "COLON"
	TokenEOF           TokenType = "EOF"
	TokenString        TokenType = "STRING"
	TokenFunction      TokenType = "FUNCTION"
	TokenReturn        TokenType = "RETURN"
	TokenIllegal       TokenType = "ILLEGAL"
	TokenComment       TokenType = "COMMENT"
	TokenDollarSign    TokenType = "DOLLAR_SIGN"
	TokenFlag          TokenType = "FLAG"
	TokenWhitespace    TokenType = "WHITESPACE"
	TokenNewline       TokenType = "NEWLINE"
	TokenSubshell      TokenType = "SUBSHELL"
)

const (
	KeywordVar      = "var"
	KeywordConst    = "const"
	KeywordEcho     = "echo"
	KeywordSource   = "source"
	KeywordIf       = "if"
	KeywordElse     = "else"
	KeywordFunc     = "func"
	KeywordReturn   = "return"
	KeywordFor      = "for"
	KeywordWhile    = "while"
	KeywordDo       = "do"
	KeywordBreak    = "break"
	KeywordContinue = "continue"
	KeywordSleep    = "sleep"
)

const (
	boolType    = "bool"
	intType     = "int"
	float64Type = "float64"
	stringType  = "string"
)

var variableKeywords = KeywordList{KeywordVar, KeywordConst}
var controlFlowKeywords = KeywordList{KeywordIf, KeywordElse, KeywordFor, KeywordWhile, KeywordDo, KeywordBreak, KeywordContinue}
var shellKeywords = KeywordList{KeywordEcho, KeywordSource, KeywordSleep}
var otherKeywords = KeywordList{KeywordFunc, KeywordReturn}
var primitiveTypes = KeywordList{boolType, intType, float64Type, stringType}

func getKeywords() KeywordList {
	return append(variableKeywords, controlFlowKeywords...)
}

func generatePattern(keywords []string) string {
	return `(\s*|^)\b(` + strings.Join(keywords, "|") + `)\b(\s*|$)`
}

func compilePattern(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

var tokenSpecs = []TokenSpec{
	{Type: TokenComment, Pattern: compilePattern(`(//.*)|(/\*[\s\S]*?\*/)`)},
	{Type: TokenKeyword, Pattern: compilePattern(generatePattern(getKeywords()))},
	{Type: TokenShellKeyword, Pattern: compilePattern(generatePattern(shellKeywords))},
	{Type: TokenPrimitiveType, Pattern: compilePattern(generatePattern(primitiveTypes))},
	{Type: TokenSubshell, Pattern: compilePattern(`\$\((.*)\)`)},

	{Type: TokenIdentifier, Pattern: compilePattern(`(\s*|^)\b[a-zA-Z_][a-zA-Z0-9_]*\b(\s*|$)`)},
	{Type: TokenFlag, Pattern: compilePattern(`-\w+`)},
	{Type: TokenNumber, Pattern: compilePattern(`\b\d+(\.\d+)?\b`)},
	{Type: TokenOperator, Pattern: compilePattern(`[+\-*/=]`)},
	{Type: TokenString, Pattern: compilePattern(`"([^\n"$])*"`)},

	{Type: TokenDollarSign, Pattern: compilePattern(`(\s*|^|)\$\w+(\s*|$)`)},
	{Type: TokenLeftParen, Pattern: compilePattern(`\(`)},
	{Type: TokenRightParen, Pattern: compilePattern(`\)`)},
	{Type: TokenLeftCurly, Pattern: compilePattern(`\{`)},
	{Type: TokenRightCurly, Pattern: compilePattern(`\}`)},
	{Type: TokenSemicolon, Pattern: compilePattern(`;`)},
	{Type: TokenNewline, Pattern: compilePattern(`\\n`)},
	{Type: TokenColon, Pattern: compilePattern(`:`)},
}

func generateTokenMap() map[TokenType]*regexp.Regexp {
	tokenMap := make(map[TokenType]*regexp.Regexp)
	for _, spec := range tokenSpecs {
		tokenMap[spec.Type] = spec.Pattern
	}
	return tokenMap
}

var tokenMap = generateTokenMap()

type Token struct {
	Type     TokenType
	Start    int
	End      int
	Value    string
	RawValue string
	Line     int
}

func (t Token) GetLine() int {
	return t.Line
}

type Lexer struct {
	Source   string
	Tokens   []Token
	Pos      int
	Filename string
	Line     int
}
