package interpreter

import (
	"os/exec"
)

func (i *Interpreter) InterpretSubShell(arguments string) string {
	cmd := exec.Command("bash", "-c", arguments)
	out, err := cmd.CombinedOutput()

	if err != nil {
		panic(string(out))
	}

	return string(out)
}
