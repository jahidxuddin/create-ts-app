package main

import (
	"os"

	"github.com/jahidxuddin/create-ts-app/internal/cli"
)

func main() {
	args := os.Args[1:]

	projectName, err := cli.ParseProjectNameFromArgs(args)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	println(projectName)
}
