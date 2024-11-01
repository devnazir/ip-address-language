package oops

import (
	"fmt"

	lx "github.com/devnazir/gosh-script/pkg/lexer"
)

var (
	createErrorMessage = func(token lx.Token, msg string, args ...interface{}) string {
		return fmt.Sprintf(msg+" at line %d", append(args, token.Line)...)
	}

	IllegalToken = func(token lx.Token) error {
		panic(New(createErrorMessage(token, "Illegal token: %s", token.Value)))
	}
	UnexpectedToken = func(token lx.Token, expected string) error {
		if expected != "" {
			panic(New(createErrorMessage(token, "Unexpected token: %s, Expected: %s", token.Value, expected)))
		}
		panic(New(createErrorMessage(token, "Unexpected token: %s", token.Value)))
	}
	UnexpectedKeyword = func(token lx.Token) error {
		panic(New(createErrorMessage(token, "Unexpected keyword: %s", token.Value)))
	}
	IllegalIdentifier = func(token lx.Token) error {
		panic(New(createErrorMessage(token, "Illegal identifier: %s", token.Value)))
	}
	ExpectedIdentifier = func(token lx.Token) error {
		panic(New(createErrorMessage(token, "Expected identifier", nil)))
	}
	ExpectedOperator = func(token lx.Token, operator string) error {
		panic(New(createErrorMessage(token, "Expected operator: %s", operator)))
	}
	TypeMismatch = func(token lx.Token, expected string, got string) error {
		panic(New(createErrorMessage(token, "Type mismatch: Expected type: %s, Got type: %s", expected, got)))
	}

	InvalidConcatenation = func(token lx.Token, operator string) error {
		panic(New(createErrorMessage(token, "Invalid concatenation %s", operator)))
	}
	ExpectedTypeAnnotation = func(token lx.Token) error {
		panic(New(createErrorMessage(token, "Expected type annotation %v", token.Value)))
	}
)
