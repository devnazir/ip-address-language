package interpreter

import (
	lx "github.com/devnazir/ip-address-language/internal/lexer"
	"github.com/devnazir/ip-address-language/pkg/ast"
	"github.com/devnazir/ip-address-language/pkg/oops"
	"github.com/devnazir/ip-address-language/pkg/semantics"
)

func (i *Interpreter) InterpretAssigmentExpression(astExpr ast.AssignmentExpression) error {
	value := i.InterpretBinaryExpr(astExpr.Expression, true)
	info := i.scopeResolver.ResolveScope(astExpr.Name)

	if info.Kind == lx.KeywordSource {
		return oops.RuntimeError(astExpr, "Cannot assign to source")
	}

	i.symbolTable.Insert(astExpr.Name, semantics.SymbolInfo{
		Kind:  lx.KeywordVar,
		Value: value,
		Line:  astExpr.Line,
	})

	return nil
}
