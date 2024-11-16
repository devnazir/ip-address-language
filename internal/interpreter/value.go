package interpreter

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/semantics"
)

type Interpreter struct {
	symbolTable   *semantics.SymbolTable
	scopeResolver *semantics.ScopeResolver
}

type IntrepretEchoStmt struct {
	expression    ast.EchoStatement
	captureOutput bool
}

type InterpretShellExpression struct {
	expression    ast.ShellExpression
	captureOutput bool
}
