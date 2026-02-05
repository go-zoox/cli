# Flags

Flags allow users to pass configuration options to your CLI application. CLI supports various flag types.

## Flag Types

### StringFlag

```go
&cli.StringFlag{
	Name:  "name",
	Usage: "Your name",
	Value: "World",
	Aliases: []string{"n"},
	EnvVars: []string{"NAME"},
}
```

### IntFlag

```go
&cli.IntFlag{
	Name:  "port",
	Usage: "Port number",
	Value: 8080,
	Aliases: []string{"p"},
	EnvVars: []string{"PORT"},
}
```

### BoolFlag

```go
&cli.BoolFlag{
	Name:  "verbose",
	Usage: "Enable verbose output",
	Value: false,
	Aliases: []string{"v"},
	EnvVars: []string{"VERBOSE"},
}
```

### StringSliceFlag

```go
&cli.StringSliceFlag{
	Name:  "tags",
	Usage: "Tags for the item",
	Aliases: []string{"t"},
	EnvVars: []string{"TAGS"},
}
```

### IntSliceFlag

```go
&cli.IntSliceFlag{
	Name:  "ports",
	Usage: "Port numbers",
	EnvVars: []string{"PORTS"},
}
```

## Flag Options

All flags support the following options:

| Option | Type | Description |
|--------|------|-------------|
| `Name` | `string` | The flag name (required) |
| `Usage` | `string` | Description of the flag |
| `Value` | `T` | Default value for the flag |
| `Aliases` | `[]string` | Short names for the flag (e.g., `-n` for `--name`) |
| `EnvVars` | `[]string` | Environment variable names that can set this flag |

## Accessing Flag Values

In your command action, access flag values using the context:

```go
app.Command(func(ctx *cli.Context) error {
	// String flags
	name := ctx.String("name")
	
	// Int flags
	port := ctx.Int("port")
	
	// Bool flags
	verbose := ctx.Bool("verbose")
	
	// Slice flags
	tags := ctx.StringSlice("tags")
	ports := ctx.IntSlice("ports")
	
	return nil
})
```

## Environment Variables

Flags can be set via environment variables:

```bash
# Set via environment variable
export NAME="John"
export PORT=3000
export VERBOSE=true

# Run the command
./myapp
```

## Examples

### Complete Example

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:  "example",
		Usage: "Example CLI with various flags",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Usage:   "Your name",
				Aliases: []string{"n"},
				EnvVars: []string{"NAME"},
				Value:   "World",
			},
			&cli.IntFlag{
				Name:    "port",
				Usage:   "Port number",
				Aliases: []string{"p"},
				EnvVars: []string{"PORT"},
				Value:   8080,
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Usage:   "Enable verbose output",
				Aliases: []string{"v"},
				EnvVars: []string{"VERBOSE"},
			},
			&cli.StringSliceFlag{
				Name:    "tags",
				Usage:   "Tags",
				Aliases: []string{"t"},
				EnvVars: []string{"TAGS"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		fmt.Printf("Name: %s\n", ctx.String("name"))
		fmt.Printf("Port: %d\n", ctx.Int("port"))
		fmt.Printf("Verbose: %t\n", ctx.Bool("verbose"))
		fmt.Printf("Tags: %v\n", ctx.StringSlice("tags"))
		return nil
	})

	app.Run()
}
```

### Usage

```bash
# Using long flags
./example --name "John" --port 3000 --verbose --tags "go,cli,example"

# Using short flags
./example -n "John" -p 3000 -v -t "go" -t "cli"

# Using environment variables
export NAME="John"
export PORT=3000
export VERBOSE=true
export TAGS="go,cli"
./example
```
