package lexer

import (
	"regexp"
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
	TokenLeftBracket   TokenType = "LEFT_BRACKET"
	TokenRightBracket  TokenType = "RIGHT_BRACKET"
	TokenSemicolon     TokenType = "SEMICOLON"
	TokenColon         TokenType = "COLON"
	TokenEOF           TokenType = "EOF"
	TokenString        TokenType = "STRING"
	TokenReturn        TokenType = "RETURN"
	TokenIllegal       TokenType = "ILLEGAL"
	TokenComment       TokenType = "COMMENT"
	TokenDollarSign    TokenType = "DOLLAR_SIGN"
	TokenFlag          TokenType = "FLAG"
	TokenWhitespace    TokenType = "WHITESPACE"
	TokenNewline       TokenType = "NEWLINE"
	TokenSubshell      TokenType = "SUBSHELL"
	TokenComma         TokenType = "COMMA"
	TokenDot           TokenType = "DOT"
	TokenBoolean       TokenType = "BOOLEAN"
	TokenArray         TokenType = "ARRAY"
	TokenObject        TokenType = "OBJECT"
	TokenTickQuote     TokenType = "TICK_QUOTE"
)

// List of keywords
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
	BoolType    = "bool"
	IntType     = "int"
	Float64Type = "float64"
	StringType  = "string"

	MultiplicationSign = "*"
	AdditionSign       = "+"
	SubtractionSign    = "-"
	DivisionSign       = "/"
	EqualsSign         = "="
	NotEqualsSign      = "!="
	GreaterThanSign    = ">"
	LessThanSign       = "<"
	GreaterOrEqualSign = ">="
	LessOrEqualSign    = "<="
	AndOperator        = "&&"
	OrOperator         = "||"
	EquivalenceSign    = "=="
	ModulusSign        = "%"

	SingleLineComment     = "//"
	MultiLineCommentStart = "/*"
	MultilineDocComment   = "/**"

	BoolTrue  = "true"
	BoolFalse = "false"
)

var Keywords = map[string]TokenType{
	KeywordVar:      TokenKeyword,
	KeywordConst:    TokenKeyword,
	KeywordEcho:     TokenShellKeyword,
	KeywordSource:   TokenShellKeyword,
	KeywordIf:       TokenKeyword,
	KeywordElse:     TokenKeyword,
	KeywordFunc:     TokenKeyword,
	KeywordReturn:   TokenKeyword,
	KeywordFor:      TokenKeyword,
	KeywordWhile:    TokenKeyword,
	KeywordDo:       TokenKeyword,
	KeywordBreak:    TokenKeyword,
	KeywordContinue: TokenKeyword,
	KeywordSleep:    TokenShellKeyword,

	// Primitive types
	BoolType:    TokenPrimitiveType,
	IntType:     TokenPrimitiveType,
	Float64Type: TokenPrimitiveType,
	StringType:  TokenPrimitiveType,

	BoolTrue:  TokenBoolean,
	BoolFalse: TokenBoolean,
}

var CommentSymbols = map[string]TokenType{
	SingleLineComment:     TokenComment,
	MultiLineCommentStart: TokenComment,
	MultilineDocComment:   TokenComment,
}

var TokenSpecs = map[TokenType]string{
	TokenSubshell:     `^\$\((.*)\)`,
	TokenDollarSign:   `^\$\{?\w+\}?`,
	TokenFlag:         `^\-[a-zA-Z]`,
	TokenNumber:       `^\b\d+(\.\d+)?\b`,
	TokenIdentifier:   `^\b[a-zA-Z_][a-zA-Z0-9_]*\b`,
	TokenOperator:     `^([+\-*/=]|==|!=|>=?|<=?|&&|\|\|)[^\w]`,
	TokenString:       `^("|')([^"\n])*("|')`,
	TokenTickQuote:    "^`",
	TokenLeftParen:    `^\(`,
	TokenRightParen:   `^\)`,
	TokenLeftCurly:    `^\{`,
	TokenRightCurly:   `^\}`,
	TokenLeftBracket:  `^\[`,
	TokenRightBracket: `^\]`,
	TokenSemicolon:    `^\;`,
	TokenColon:        `^:`,
	TokenDot:          `^\.`,
	TokenComma:        `^,`,
	TokenNewline:      `^\\n`,
}

type Token struct {
	Type     TokenType
	Start    int
	End      int
	Value    string
	RawValue string
	Line     int
	Index    int
}

func (t Token) GetLine() int {
	return t.Line
}

func (t Token) GetType() interface{} {
	return t.Type
}

type Lexer struct {
	Source       string
	Tokens       []Token
	Pos          int
	Filename     string
	Line         int
	CurrentIndex int
}
