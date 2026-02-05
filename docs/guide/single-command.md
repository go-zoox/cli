# Single Command

Single command CLI applications are perfect for simple tools that perform a single action. This is the simplest way to create a CLI application.

## Basic Example

```go
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
			&cli.StringFlag{
				Name:  "string",
				Usage: "string flag",
				Value: "default value",
			},
			&cli.IntFlag{
				Name:  "int",
				Usage: "int flag",
				Value: 0,
			},
			&cli.BoolFlag{
				Name:  "bool",
				Usage: "bool flag",
				Value: false,
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		fmt.Printf("String: %s\n", ctx.String("string"))
		fmt.Printf("Int: %d\n", ctx.Int("int"))
		fmt.Printf("Bool: %t\n", ctx.Bool("bool"))
		return nil
	})

	app.Run()
}
```

## Configuration Options

### SingleProgramConfig

| Field | Type | Description |
|-------|------|-------------|
| `Name` | `string` | The name of the program |
| `HelpName` | `string` | Full name of command for help |
| `Usage` | `string` | Description of the program |
| `UsageText` | `string` | Text to override the USAGE section of help |
| `ArgsUsage` | `string` | Description of the program argument format |
| `Version` | `string` | Version of the program |
| `Description` | `string` | Description of the program |
| `EnableBashCompletion` | `bool` | Boolean to enable bash completion commands |
| `HideHelp` | `bool` | Boolean to hide built-in help command and help flag |
| `HideHelpCommand` | `bool` | Boolean to hide built-in help command but keep help flag |
| `HideVersion` | `bool` | Boolean to hide built-in version flag |
| `Flags` | `[]Flag` | List of flags to parse |

## Usage

```bash
# Run with default values
./single

# Run with flags
./single --string "hello" --int 42 --bool

# Show help
./single --help

# Show version
./single --version
```
