// Package shell is for executing commands in the user terminal
package shell

import (
	"os"
	"os/exec"
)

func RunCommand(showOutput bool, name string, args ...string) error {
	cmd := exec.Command(name, args...)

	cmd.Stdin = os.Stdin

	if showOutput {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
}
