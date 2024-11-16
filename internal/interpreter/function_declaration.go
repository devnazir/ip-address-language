package interpreter

import (
	"fmt"
	"runtime/debug"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (i *Interpreter) InterpretBodyFunction(p ast.FunctionDeclaration) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
		}
	}()

	i.scopeResolver.EnterScope()
	for _, nodeItem := range p.Body {
		InterpretNode(i, nodeItem, "")
	}
	i.scopeResolver.ExitScope()
}
