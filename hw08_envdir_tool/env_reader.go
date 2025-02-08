package main

import (
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := os.ReadFile(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		lines := strings.Split(string(content), "\n")
		value := strings.TrimRight(lines[0], " \t\r\n")
		value = strings.ReplaceAll(value, "\x00", "\n")

		env[file.Name()] = EnvValue{
			Value:      value,
			NeedRemove: len(value) == 0,
		}
	}

	return env, nil
}
