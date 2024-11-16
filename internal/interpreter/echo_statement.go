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

			info := i.scopeResolver.ResolveScope(identifier.Name)
			value := info.Value

			if reflect.TypeOf(value).Kind() == reflect.Int {
				value = strconv.Itoa(value.(int))
			}

			vars := utils.FindShellVars(value.(string))
			for _, v := range vars {
				info := i.scopeResolver.ResolveScope(v[1:])
				value = strings.ReplaceAll(value.(string), v, info.Value.(string))
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
			value := literal.Value

			vars := utils.FindShellVars(value)
			for _, v := range vars {
				info := i.scopeResolver.ResolveScope(v[1:])
				if _, ok := info.Value.(int); ok {
					value = strings.ReplaceAll(value, v, strconv.Itoa(info.Value.(int)))
					continue
				}

				value = strings.ReplaceAll(value, v, info.Value.(string))
			}

			if literal.Value == "echo" {
				cmdArgs += fmt.Sprintf("%v", value) + " -e '\n'"
				break
			}

			cmdArgs += fmt.Sprintf("%v", value)

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
