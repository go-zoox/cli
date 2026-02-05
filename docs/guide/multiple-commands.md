# Multiple Commands

Multiple commands CLI applications allow you to create a CLI tool with multiple subcommands, similar to `git` or `docker`.

## Basic Example

```go
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
			&cli.StringFlag{
				Name:  "format",
				Usage: "Output format",
				Value: "table",
			},
		},
		Action: func(ctx *cli.Context) error {
			format := ctx.String("format")
			fmt.Printf("Listing items in %s format\n", format)
			return nil
		},
	})

	app.Register("create", &cli.Command{
		Name:  "create",
		Usage: "create is a command that creates things.",
		Action: func(ctx *cli.Context) error {
			fmt.Println("Creating item...")
			return nil
		},
	})

	app.Register("delete", &cli.Command{
		Name:  "delete",
		Usage: "delete is a command that deletes things.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "force",
				Usage: "Force deletion",
				Value: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			force := ctx.Bool("force")
			if force {
				fmt.Println("Force deleting item...")
			} else {
				fmt.Println("Deleting item...")
			}
			return nil
		},
	})

	app.Run()
}
```

## Usage

```bash
# Show all commands
./multiple --help

# Run list command
./multiple list

# Run list with flags
./multiple list --format json

# Run create command
./multiple create

# Run delete command
./multiple delete --force
```

## Command Structure

Each command can have:

- **Name**: The command name
- **Usage**: Description of what the command does
- **Flags**: Command-specific flags
- **Action**: The function to execute when the command is called

## Nested Commands

You can create nested commands by using dots in the command name:

```go
app.Register("user.create", &cli.Command{
	Name:  "user.create",
	Usage: "Create a new user",
	Action: func(ctx *cli.Context) error {
		// ...
		return nil
	},
})
```

Usage:

```bash
./multiple user.create
```
