package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, envValue := range env {
		if envValue.NeedRemove {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, envValue.Value)
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		return 1
	}

	return 0
}
