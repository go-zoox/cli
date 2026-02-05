# Installation

## Prerequisites

- Go 1.18 or higher

## Install

Install the package using `go get`:

```bash
go get github.com/go-zoox/cli
```

Or with a specific version:

```bash
go get github.com/go-zoox/cli@latest
```

## Verify Installation

Create a simple test file to verify the installation:

```go
package main

import (
	"fmt"
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "test",
		Usage:   "Test CLI installation",
		Version: "1.0.0",
	})

	app.Command(func(ctx *cli.Context) error {
		fmt.Println("CLI is installed correctly!")
		return nil
	})

	app.Run()
}
```

Run it:

```bash
go run main.go
# Output: CLI is installed correctly!
```

## Import in Your Project

```go
import "github.com/go-zoox/cli"
```

For interactive components:

```go
import "github.com/go-zoox/cli/interactive"
```

For loading components:

```go
import "github.com/go-zoox/cli/loading"
```
