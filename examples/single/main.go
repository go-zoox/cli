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
	})

	app.Command(func(ctx *cli.Context) error {
		fmt.Println("i am a single")
		return nil
	})

	app.Run()
}
