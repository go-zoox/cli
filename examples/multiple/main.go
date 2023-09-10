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
		Flags: []cli.Flag{
			// String
			&cli.StringFlag{
				Name:  "string",
				Usage: "string flag",
				// Environment
				EnvVars: []string{"ENV_VAR"},
				// Short name
				Aliases: []string{"s"},
				Value:   "default value",
			},

			// Int
			&cli.IntFlag{
				Name:  "int",
				Usage: "int flag",
				// Environment
				EnvVars: []string{"ENV_VAR"},
				// Short name
				Aliases: []string{"i"},
				Value:   0,
			},

			// Bool
			&cli.BoolFlag{
				Name:  "bool",
				Usage: "bool flag",
				// Environment
				EnvVars: []string{"ENV_VAR"},
				// Short name
				Aliases: []string{"b"},
				Value:   false,
			},

			// String Array
			&cli.StringSliceFlag{
				Name:  "string-array",
				Usage: "string array flag",
				// Environment
				EnvVars: []string{"ENV_VAR"},
				// Short name
				Aliases: []string{"x"},
			},
		},
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
