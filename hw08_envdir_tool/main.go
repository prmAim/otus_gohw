package main

import "os"

func main() {
	if len(os.Args) < 3 {
		panic("Usage: go-envdir <env_dir> <command> [args...]")
	}

	envDir := os.Args[1]
	cmd := os.Args[2:]

	env, err := ReadDir(envDir)
	if err != nil {
		panic(err)
	}

	os.Exit(RunCmd(cmd, env))
}
