package interpreter

import (
	"fmt"
	"os/exec"
)

func (i *Interpreter) InterpretLsStmt(params InterpretLsStmt) string {
	captureOutput := params.captureOutput

	command := "ls"

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
