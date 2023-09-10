package main

import (
	"fmt"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "single",
		Usage:   "single is a program that has a single command.",
		Version: "0.0.1",
	})

	app.Command(func(ctx *cli.Context) error {
		fmt.Println("i am a single")
		return nil
	})

	app.Serve()
}
