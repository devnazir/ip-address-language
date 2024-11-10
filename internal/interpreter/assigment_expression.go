package interpreter

import (
	"strings"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (i *Interpreter) InterpretAssigmentExpression(astExpr ast.AssignmentExpression) {
	name := strings.TrimSpace(astExpr.Name)
	value := i.InterpretBinaryExpr(astExpr.Expression)
	env.SetVariable(name, value)
}
