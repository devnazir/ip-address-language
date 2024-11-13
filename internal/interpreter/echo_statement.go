package interpreter

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/utils"
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

			vars := utils.FindShellVars(value.(string))
			for _, v := range vars {
				value = strings.ReplaceAll(value.(string), v, env.GetVariable(v[1:]).(string))
			}

			cmdArgs += fmt.Sprintf("%v", value) + " "
		case ast.NumberLiteral:
			literal := argument.(ast.NumberLiteral)

			if reflect.TypeOf(literal.Value).Kind() == reflect.Int {
				literal.Value = strconv.Itoa(literal.Value.(int))
			}

			cmdArgs += fmt.Sprintf("%v", literal.Raw)
		case ast.StringLiteral:
			literal := argument.(ast.StringLiteral)

			if literal.Value == "echo" {
				cmdArgs += fmt.Sprintf("%v", literal.Value) + " -e '\n'"
				break
			}

			cmdArgs += fmt.Sprintf("%v", literal.Raw)

		case ast.SubShell:
			subShell := argument.(ast.SubShell)
			cmdArgs += "'$(" + fmt.Sprintf("%v", subShell.Arguments) + ")'"

		case ast.Illegal:
			illegal := argument.(ast.Illegal)
			cmdArgs += fmt.Sprintf("%v", illegal.Value) + " "

		default:
			panic("Invalid argument type")
		}
	}

	command, _ := utils.RemoveDoubleQuotes(cmdFlags + "'" + cmdArgs + "'")
	cmd := exec.Command("bash", "-c", "echo "+command)
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic(string(out))
	}

	if captureOutput {
		return fmt.Sprintf("%s", out)
	}

	fmt.Printf("%s", out)
	return ""
}
