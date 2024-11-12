package interpreter

import (
	"strings"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

func (i *Interpreter) InterpretVariableDeclaration(nodeVar ast.VariableDeclaration) {
	name := strings.TrimSpace(nodeVar.Declarations[0].Id.(ast.Identifier).Name)

	if env.HasVariable(name) {
		oops.DuplicateIdentifierError(nodeVar)
	}

	if _, ok := nodeVar.Declarations[0].Init.(ast.SubShell); ok {
		res := i.InterpretSubShell(nodeVar.Declarations[0].Init.(ast.SubShell).Arguments.(string))
		env.SetVariable(name, res)
		return
	}

	value := i.InterpretBinaryExpr(nodeVar.Declarations[0].Init)
	env.SetVariable(name, value)
}
