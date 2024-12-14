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
	TokenDoubleQuote   TokenType = "DOUBLE_QUOTE"
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
	LeftParen          = "("
	RightParen         = ")"
	Comma              = ","
	Semicolon          = ";"
	Colon              = ":"
	Dot                = "."
	LeftCurly          = "{"
	RightCurly         = "}"
	LeftBracket        = "["
	RightBracket       = "]"
	TickQuote          = "`"
	DoubleQuote        = "\""
	Newline            = "\\n"
	DollarSign         = "$"

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

	LeftParen:          TokenLeftParen,
	RightParen:         TokenRightParen,
	Comma:              TokenComma,
	Semicolon:          TokenSemicolon,
	Colon:              TokenColon,
	Dot:                TokenDot,
	LeftCurly:          TokenLeftCurly,
	RightCurly:         TokenRightCurly,
	LeftBracket:        TokenLeftBracket,
	RightBracket:       TokenRightBracket,
	TickQuote:          TokenTickQuote,
	DoubleQuote:        TokenDoubleQuote,
	Newline:            TokenNewline,
	ModulusSign:        TokenOperator,
	MultiplicationSign: TokenOperator,
	AdditionSign:       TokenOperator,
	SubtractionSign:    TokenOperator,
	DivisionSign:       TokenOperator,
	EqualsSign:         TokenOperator,
	NotEqualsSign:      TokenOperator,
	GreaterThanSign:    TokenOperator,
	LessThanSign:       TokenOperator,
	GreaterOrEqualSign: TokenOperator,
	LessOrEqualSign:    TokenOperator,
	AndOperator:        TokenOperator,
	OrOperator:         TokenOperator,
	EquivalenceSign:    TokenOperator,
	DollarSign:         TokenDollarSign,

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

const (
	TokenFlagRegex       = `^\-[a-zA-Z]`
	TokenSubshellRegex   = `^\$\((.*)\)`
	TokenDollarSignRegex = `^\$\{?\w+\}?`
)

var TokenSpecs = map[TokenType]string{
	TokenSubshell:   TokenSubshellRegex,
	TokenDollarSign: TokenDollarSignRegex,
	TokenFlag:       TokenFlagRegex,
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
