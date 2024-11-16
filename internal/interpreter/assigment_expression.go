package interpreter

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/semantics"
)

func (i *Interpreter) InterpretAssigmentExpression(astExpr ast.AssignmentExpression) {
	value := i.InterpretBinaryExpr(astExpr.Expression)
	i.symbolTable.Insert(astExpr.Name, semantics.SymbolInfo{
		Type:  "",
		Value: value,
	})
}
