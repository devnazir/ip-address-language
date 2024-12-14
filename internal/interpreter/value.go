package interpreter

import (
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/semantics"
)

type Interpreter struct {
	symbolTable   *semantics.SymbolTable
	scopeResolver *semantics.ScopeResolver
}

type IntrepretEchoStmt struct {
	expression ast.EchoStatement
}

type InterpretShellExpression struct {
	expression    ast.ShellExpression
	captureOutput bool
}
