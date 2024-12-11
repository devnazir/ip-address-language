package interpreter

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"
	"strings"

	"github.com/devnazir/gosh-script/pkg/ast"
	"github.com/devnazir/gosh-script/pkg/oops"
	"github.com/devnazir/gosh-script/pkg/utils"
)

func (i *Interpreter) IntrepretEchoStmt(params IntrepretEchoStmt) string {
	expression := params.expression

	echoArguments := expression.Arguments
	echoFlags := expression.Flags

	var cmdArgs string
	var cmdFlags string

	for _, flag := range echoFlags {
		cmdFlags += flag + " "
	}

	for _, argument := range echoArguments {
		cmdArgs += i.processArgument(argument)
	}

	command, _ := utils.RemoveDoubleQuotes(cmdFlags + "'" + cmdArgs + "'")
	cmd := exec.Command("bash", "-c", "echo "+command)
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic(string(out))
	}

	return fmt.Sprintf("%s", out)
}

func (i *Interpreter) processArgument(argument ast.ASTNode) string {
	var result string

	switch argument.GetType() {
	case ast.IdentifierTree:
		identifier := argument.(ast.Identifier)
		info := i.scopeResolver.ResolveScope(identifier.Name)
		value := i.resolveValue(info.Value)
		result = formatValue(value)
	case ast.NumberLiteralTree:
		literal := argument.(ast.NumberLiteral)
		result = formatValue(literal.Value)
	case ast.StringLiteralTree:
		literal := argument.(ast.StringLiteral)
		result = i.resolveStringLiteral(literal)
	case ast.StringTemplateLiteralTree:
		literal := argument.(ast.StringTemplateLiteral)
		for _, part := range literal.Parts {
			result += i.processArgument(part)
		}
	case ast.SubShellTree:
		subShell := argument.(ast.SubShell)
		result = fmt.Sprintf("'$(%v)'", subShell.Arguments)
	case ast.IllegalTree:
		illegal := argument.(ast.Illegal)
		result = illegal.Value
	case ast.MemberExpressionTree:
		memberExpr := argument.(ast.MemberExpression)
		result = fmt.Sprintf("%v ", i.InterpretMemberExpr(memberExpr))
	default:
		oops.InvalidEchoArgumentError(argument.(ast.EchoStatement))
	}

	return result
}

func (i *Interpreter) resolveValue(value interface{}) string {
	if reflect.TypeOf(value).Kind() == reflect.Int {
		return strconv.Itoa(value.(int))
	}
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		var result string
		for _, v := range value.([]interface{}) {
			result += fmt.Sprintf("%v", v)
		}
		return result
	}
	return fmt.Sprintf("%v", value)
}

func (i *Interpreter) resolveStringLiteral(literal ast.StringLiteral) string {
	value := literal.Raw
	vars := utils.FindShellVars(value)

	for _, v := range vars {
		info := i.scopeResolver.ResolveScope(v[1:])
		resolvedValue := i.resolveValue(info.Value)
		value = strings.ReplaceAll(value, v, resolvedValue)
	}

	if literal.Value == "echo" {
		return fmt.Sprintf("%v -e '\n'", value)
	}

	return value
}

func formatValue(value interface{}) string {
	return fmt.Sprintf("%v", value)
}
