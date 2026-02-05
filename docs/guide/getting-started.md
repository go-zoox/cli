# Getting Started

CLI is a simple and powerful command line framework for Go. It provides an easy-to-use API for building CLI applications with support for single commands, multiple commands, interactive components, and loading indicators.

## Quick Start

### Installation

```bash
go get github.com/go-zoox/cli
```

### Your First CLI Application

Here's a simple example of a single command CLI:

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "myapp",
		Usage:   "A simple CLI application",
		Version: "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your name",
				Value: "World",
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		name := ctx.String("name")
		fmt.Printf("Hello, %s!\n", name)
		return nil
	})

	app.Run()
}
```

Run it:

```bash
go run main.go --name "CLI"
# Output: Hello, CLI!
```

## What's Next?

- Learn about [Single Command](/guide/single-command) applications
- Explore [Multiple Commands](/guide/multiple-commands) support
- Check out [Interactive Components](/guide/interactive/text) for better UX
- See [Examples](/examples/) for more use cases
