package oops

import (
	"fmt"

	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
)

var (
	IllegalTokenError = func(illegalToken lx.Token) error {
		panic(New(CreateErrorMessage(illegalToken, "Illegal token: %s", illegalToken.Value)))
	}

	UnexpectedTokenError = func(unexpectedToken lx.Token, expectedToken string) error {
		if expectedToken != "" {
			panic(New(CreateErrorMessage(unexpectedToken, "Unexpected token: %s, Expected: %s", unexpectedToken.Value, expectedToken)))
		}
		panic(New(CreateErrorMessage(unexpectedToken, "Unexpected token: %s", unexpectedToken.Value)))
	}

	UnexpectedKeywordError = func(unexpectedKeyword lx.Token) error {
		panic(New(CreateErrorMessage(unexpectedKeyword, "Unexpected keyword: %s", unexpectedKeyword.Value)))
	}

	IllegalIdentifierError = func(illegalIdentifier lx.Token) error {
		panic(New(CreateErrorMessage(illegalIdentifier, "Illegal identifier: %s", illegalIdentifier.Value)))
	}

	ExpectedIdentifierError = func(missingIdentifier lx.Token) error {
		panic(New(CreateErrorMessage(missingIdentifier, "Expected identifier")))
	}

	ExpectedOperatorError = func(missingOperatorToken lx.Token, expectedOperator string) error {
		panic(New(CreateErrorMessage(missingOperatorToken, "Expected operator: %s", expectedOperator)))
	}

	TypeMismatchError = func(mismatchToken lx.Token, expectedType string, receivedType string) error {
		panic(New(CreateErrorMessage(mismatchToken, "Type mismatch: Expected type: %s, Got type: %s", expectedType, receivedType)))
	}

	InvalidConcatenationError = func(concatToken lx.Token, invalidOperator string) error {
		panic(New(CreateErrorMessage(concatToken, "Invalid concatenation with operator %s", invalidOperator)))
	}

	ExpectedTypeAnnotationError = func(identifierToken lx.Token) error {
		panic(New(CreateErrorMessage(identifierToken, "Expected type annotation for identifier %v", identifierToken.Value)))
	}

	ExpectedEntrypointFileError = func() error {
		panic(New("Expected filename as main entrypoint, e.g., gsh main.gsh"))
	}

	ExpectedTokenError = func(currentToken lx.Token, expectedToken string) error {
		panic(New(CreateErrorMessage(currentToken, "Expected token: %s", expectedToken)))
	}

	DuplicateIdentifierError = func(variableDecl ast.VariableDeclaration) error {
		panic(New(CreateErrorMessage(variableDecl, "Identifier '%s' has already been declared", variableDecl.Declaration.Id.(ast.Identifier).Name)))
	}

	IdentifierNotFoundError = func(identifierToken ast.Identifier) error {
		panic(New(CreateErrorMessage(identifierToken, "Identifier '%s' not found", identifierToken.Name)))
	}

	InvalidFileExtensionError = func(filename string) error {
		panic(fmt.Errorf("Invalid file extension: %s, expected .n or .nnn", filename))
	}

	// SourceAliasCannotBeAssignedError = func(symbolInfo *semantics.SymbolInfo) error {
	// 	panic(CreateErrorMessage(symbolInfo, "Source alias cannot be assigned to a variable"))
	// }

	SourceAliasMustBeAlphanumericError = func(alias string) error {
		panic(fmt.Errorf("Source alias must be alphanumeric, got: %s", alias))
	}

	InitFunctionCannotHaveParametersError = func(fNode ast.FunctionDeclaration) error {
		panic(New(CreateErrorMessage(fNode, "init function cannot have parameters")))
	}

	InvalidEchoArgumentError = func(echoNode ast.EchoStatement) error {
		panic(New(CreateErrorMessage(echoNode, "Invalid argument for echo statement")))
	}

	FunctionNotCalledWithEnoughArgumentsError = func(fNode ast.FunctionDeclaration, expectedArgs int, receivedArgs int) error {
		panic(New(CreateErrorMessage(fNode, "Function '%s' expects %d arguments, got %d", fNode.Identifier.Name, expectedArgs, receivedArgs)))
	}

	FunctionCalledWithTooManyArgumentsError = func(fNode ast.FunctionDeclaration, expectedArgs int, receivedArgs int) error {
		panic(New(CreateErrorMessage(fNode, "Function '%s' expects %d arguments, got %d", fNode.Identifier.Name, expectedArgs, receivedArgs)))
	}

	RestParameterMustBeLastError = func(fNode ast.FunctionDeclaration) error {
		panic(New(CreateErrorMessage(fNode, "Rest parameter must be the last parameter")))
	}
)
