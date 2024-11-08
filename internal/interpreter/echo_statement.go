package interpreter

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"

	"github.com/devnazir/gosh-script/pkg/ast"
)

func (i *Interpreter) IntrepretEchoStmt(params IntrepretEchoStmt) string {
	expression := params.expression
	captureOutput := params.captureOutput

	echoArguments := expression.Arguments
	echoFlags := expression.Flags

	var cmdArgs string
	var cmdFlags string

	for _, flag := range echoFlags {
		cmdFlags += flag + " "
	}

	for _, argument := range echoArguments {
		switch argument.(type) {
		case ast.Identifier:
			identifier := argument.(ast.Identifier)
			value := env.GetVariable(identifier.Name)

			if reflect.TypeOf(value).Kind() == reflect.Int {
				value = strconv.Itoa(value.(int))
			}

			cmdArgs += fmt.Sprintf("%v", value) + " "

		case ast.Literal:
			literal := argument.(ast.Literal)

			if reflect.TypeOf(literal.Value).Kind() == reflect.Int {
				literal.Value = strconv.Itoa(literal.Value.(int))
			}

			cmdArgs += fmt.Sprintf("%v", literal.Value) + " "
		}
	}

	command := "echo " + cmdFlags + "'" + cmdArgs + "'"
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic("Error executing command:" + err.Error())
	}

	if captureOutput {
		return fmt.Sprintf("%s", out)
	}

	fmt.Printf("%s", out)
	return ""
}
