package interpreter

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"

	"github.com/devnazir/gosh-script/pkg/parser"
)

var env = NewEnvironment()

func NewInterpreter() *Interpreter {
	return &Interpreter{Environment: env}
}

func (i *Interpreter) Interpret(p parser.ASTNode) {
	program := p.(parser.Program)

	for _, node := range program.Body {
		switch node.(type) {
		case parser.VariableDeclaration:
			i.InterpretVariableDeclaration(node.(parser.VariableDeclaration))

		case parser.ShellExpression:
			shell := node.(parser.ShellExpression)
			expression := shell.Expression

			switch expression.(type) {
			case parser.EchoStatement:
				i.IntrepretEchoStmt(expression.(parser.EchoStatement))
			}
		}
	}
}

func (i *Interpreter) InterpretVariableDeclaration(node parser.VariableDeclaration) {
	name := node.Declarations[0].Id.(parser.Identifier).Name
	value := node.Declarations[0].Init.(parser.Literal).Value

	if env.HasVariable(name) {
		// TODO: create err message
		panic("Duplicate Variable")
	}

	env.SetVariable(name, value)
}

func (i *Interpreter) IntrepretEchoStmt(expression parser.EchoStatement) {
	echoArguments := expression.Arguments
	echoFlags := expression.Flags

	var cmdArgs string
	var cmdFlags string

	for _, flag := range echoFlags {
		cmdFlags += flag + " "
	}

	for _, argument := range echoArguments {
		switch argument.(type) {
		case parser.Identifier:
			identifier := argument.(parser.Identifier)
			value := env.GetVariable(identifier.Name)

			if reflect.TypeOf(value).Kind() == reflect.Int {
				value = strconv.Itoa(value.(int))
			}

			cmdArgs += value.(string) + " "
		case parser.Literal:
			literal := argument.(parser.Literal)
			cmdArgs += literal.Value.(string) + " "
		}
	}

	command := "echo " + cmdFlags + "'" + cmdArgs + "'"
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic("Error executing command:" + err.Error())
	}

	fmt.Printf("%s", out)
}
