package main

import (
	"fmt"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:  "multiple",
		Usage: "multiple is a program that has multiple commands.",
	})

	app.Register("list", &cli.Command{
		Name:  "list",
		Usage: "list is a command that lists things.",
		Action: func(ctx *cli.Context) error {
			fmt.Println("i am a list")
			return nil
		},
	})

	app.Register("create", &cli.Command{
		Name:  "create",
		Usage: "create is a command that creates things.",
		Action: func(ctx *cli.Context) error {
			fmt.Println("i am a create")
			return nil
		},
	})

	app.Serve()
}
