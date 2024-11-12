package interpreter

import (
	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
)

type Interpreter struct {
	Environment *Environment
}

type Variables []map[string]interface{}

type Environment struct {
	Variables Variables
}

func NewEnvironment() *Environment {
	return &Environment{Variables: []map[string]interface{}{
		make(map[string]interface{}),
	}}
}

func (e *Environment) SetVariable(name string, value interface{}) {
	e.Variables[0][name] = value
}

func (e *Environment) GetVariable(name string) interface{} {
	for _, variable := range e.Variables {
		if val, ok := variable[name]; ok {
			return val
		}
	}

	panic(oops.IdentifierNotFoundError(ast.Identifier{Name: name}))
}

func (e *Environment) HasVariable(name string) bool {
	for _, variable := range e.Variables {
		if _, ok := variable[name]; ok {
			return true
		}
	}

	return false
}

func (e *Environment) AddScope() {
	e.Variables = append([]map[string]interface{}{{}}, e.Variables...)
}

func (e *Environment) CloseScope() {
	e.Variables = e.Variables[1:]
}

type IntrepretEchoStmt struct {
	expression    ast.EchoStatement
	captureOutput bool
}

type InterpretShellExpression struct {
	expression    ast.ShellExpression
	captureOutput bool
}
