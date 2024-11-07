package interpreter

import (
	"fmt"
	"os/exec"
	"reflect"
	"strconv"

	"github.com/devnazir/gosh-script/pkg/node"
	"github.com/devnazir/gosh-script/pkg/oops"
)

var env = NewEnvironment()

func NewInterpreter() *Interpreter {
	return &Interpreter{Environment: env}
}

func (i *Interpreter) Interpret(p node.ASTNode) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	program := p.(node.Program)

	for _, nodeItem := range program.Body {
		switch nodeItem.(type) {
		case node.VariableDeclaration:
			i.InterpretVariableDeclaration(nodeItem.(node.VariableDeclaration))

		case node.ShellExpression:
			shell := nodeItem.(node.ShellExpression)
			expression := shell.Expression

			switch expression.(type) {
			case node.EchoStatement:
				i.IntrepretEchoStmt(expression.(node.EchoStatement))
			}
		}
	}
}

func (i *Interpreter) InterpretVariableDeclaration(nodeVar node.VariableDeclaration) {
	name := nodeVar.Declarations[0].Id.(node.Identifier).Name

	if env.HasVariable(name) {
		oops.DuplicateIdentifierError(nodeVar)
	}

	value := i.InterpretBinaryExpr(nodeVar.Declarations[0].Init)
	env.SetVariable(name, value)
}

func (i *Interpreter) IntrepretEchoStmt(expression node.EchoStatement) {
	echoArguments := expression.Arguments
	echoFlags := expression.Flags

	var cmdArgs string
	var cmdFlags string

	for _, flag := range echoFlags {
		cmdFlags += flag + " "
	}

	for _, argument := range echoArguments {
		switch argument.(type) {
		case node.Identifier:
			identifier := argument.(node.Identifier)
			value := env.GetVariable(identifier.Name)

			if reflect.TypeOf(value).Kind() == reflect.Int {
				value = strconv.Itoa(value.(int))
			}

			cmdArgs += fmt.Sprintf("%v", value) + " "

		case node.Literal:
			literal := argument.(node.Literal)

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

	fmt.Printf("%s", out)
}

func (i *Interpreter) InterpretBinaryExpr(b node.ASTNode) interface{} {
	if reflect.TypeOf(b) == reflect.TypeOf(node.Literal{}) {
		return b.(node.Literal).Value
	}

	if reflect.TypeOf(b) == reflect.TypeOf(node.Identifier{}) {
		name := b.(node.Identifier).Name
		value := env.GetVariable(name)

		if value == nil {
			oops.IdentifierNotFoundError(b.(node.Identifier))
		}

		return value
	}

	if reflect.TypeOf(b) == reflect.TypeOf(node.BinaryExpression{}) {
		leftValue := i.InterpretBinaryExpr(b.(node.BinaryExpression).Left)
		rightValue := i.InterpretBinaryExpr(b.(node.BinaryExpression).Right)
		operator := b.(node.BinaryExpression).Operator

		var leftFloat, rightFloat float64
		var isLeftFloat, isRightFloat bool

		switch v := leftValue.(type) {
		case int:
			leftFloat = float64(v)
		case float64:
			leftFloat = v
			isLeftFloat = true
		default:
			return 0
		}

		switch v := rightValue.(type) {
		case int:
			rightFloat = float64(v)
		case float64:
			rightFloat = v
			isRightFloat = true
		default:
			return 0
		}

		var result interface{}
		switch operator {
		case "+":
			result = leftFloat + rightFloat
		case "-":
			result = leftFloat - rightFloat
		case "*":
			result = leftFloat * rightFloat
		case "/":
			if rightFloat == 0 {
				return 0
			}
			result = leftFloat / rightFloat
		}

		if isLeftFloat || isRightFloat {
			return result
		}
		return int(result.(float64))
	}

	return 0
}
