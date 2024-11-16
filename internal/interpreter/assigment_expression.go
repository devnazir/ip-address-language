package interpreter

import (
	lx "github.com/devnazir/gosh-script/internal/lexer"
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/semantics"
)

func (i *Interpreter) InterpretAssigmentExpression(astExpr ast.AssignmentExpression) {
	value := i.InterpretBinaryExpr(astExpr.Expression)
	info := i.scopeResolver.ResolveScope(astExpr.Name)

	if info.Kind == lx.KeywordSource {
		oops.SourceAliasCannotBeAssignedError(info)
	}

	i.symbolTable.Insert(astExpr.Name, semantics.SymbolInfo{
		Kind:  lx.KeywordVar,
		Value: value,
		Line:  astExpr.Line,
	})
}
