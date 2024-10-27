package oops

import (
	"strconv"

	lx "github.com/devnazir/gosh-script/pkg/lexer"
)

var (
	IllegalToken = func(token lx.Token) error {
		panic(New("Illegal token: " + token.Value + " at line " + strconv.Itoa(token.Line)))
	}
	UnexpectedToken = func(token lx.Token) error {
		panic(New("Unexpected token: " + token.Value + " at line " + strconv.Itoa(token.Line)))
	}
	UnexpectedKeyword = func(token lx.Token) error {
		panic(New("Unexpected keyword: " + token.Value + " at line " + strconv.Itoa(token.Line)))
	}
	IllegalIdentifier = func(token lx.Token) error {
		panic(New("Illegal identifier: " + token.Value + " at line " + strconv.Itoa(token.Line)))
	}
	ExpectedIdentifier = func(token lx.Token) error {
		panic(New("Expected identifier, error at line " + strconv.Itoa(token.Line)))
	}
	ExpectedOperator = func(token lx.Token, operator string) error {
		panic(New("Expected operator " + operator + " error at line " + strconv.Itoa(token.Line)))
	}
)
