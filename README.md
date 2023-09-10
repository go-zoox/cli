# CLI - A simple and powerful command line framework for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/cli)](https://pkg.go.dev/github.com/go-zoox/cli)
[![Build Status](https://github.com/go-zoox/cli/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/cli/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/cli)](https://goreportcard.com/report/github.com/go-zoox/cli)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/cli/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/cli?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/cli.svg)](https://github.com/go-zoox/cli/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/cli.svg?label=Release)](https://github.com/go-zoox/cli/releases)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/cli
```

## Getting Started

### Example: Single Command CLI

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
```

### Example: Multi Command CLI

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

	app.Run()
}
```

### Example: Interactive CLI - Text

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
  // Component: Text
	name, err := interactive.Text("What is your name ?", &interactive.TextOptions{
		Default:  "Zero",
		Required: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Your name is：", name)
}
```

### Example: Interactive CLI - Select

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	name, err := interactive.Select(
		"What is your favorite city?",
		[]interactive.SelectOption{
			{
				Label: "Beijing",
				Value: "beijing",
			},
			{
				Label: "Shanghai",
				Value: "shanghai",
			},
			{
				Label: "Guangzhou",
				Value: "guangzhou",
			},
		},
		&interactive.SelectOptions{
			Default: "guangzhou",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your favorite city: ", name)
}
```

### Example: Interactive CLI - Confirm

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	like, err := interactive.Confirm(
		"Do you like the book ?",
		&interactive.ConfirmOptions{
			Default: false,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your answer: ", like)
}
```

### Example: Interactive CLI - Password

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	name, err := interactive.Password("Please type your password ?", &interactive.PasswordOptions{
		Required: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Your password is：", name)
}
```

### Example: Loading - Spinner

```go
package main

import (
	"time"

	"github.com/go-zoox/cli/loading"
)

func main() {
	bar := loading.Spinner("Loading...")
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}

	bar.Finish()
}
```

### Example: Loading - Progress

```go
package main

import (
	"time"

	"github.com/go-zoox/cli/loading"
)

func main() {
	bar := loading.Progress(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}
```


## Who is using it?

- [whatwewant/chatgpt-for-chatbot-feishu](https://github.com/whatwewant/chatgpt-for-chatbot-feishu) - 快速将 ChatGPT 接入飞书，基于 OpenAI 官方接口，作为私人工作助理或者企业员工助理.
- [go-zoox/serve](https://github.com/go-zoox/serve) - A HTTP Static Server for Frontend, make you works with SPA easier.
- [go-zoox/terminal](https://github.com/go-zoox/terminal) - A Web Terminal Server / Client | Agent / Registry.
- [go-zoox/connect](https://github.com/go-zoox/connect) - Make Auth Connect Easier. Support OAuth2 and OIDC Providers. Built in support doreamon.
- [go-zoox/gzfly](https://github.com/go-zoox/gzfly) - Make Tunnel Easier Like V2Fly + Clash. Custom Protocol, based on WebSocket.
- [go-zoox/gzproxy](https://github.com/go-zoox/gzproxy) - Easy to proxy with your http server or any another upstream. Built in supports Basic Auth, Bearer Toke, OAuth2 (GitHub, Feishu, Doreamon, etc.) .
- [go-zoox/gzcaas](https://github.com/go-zoox/gzcaas) - CLI for CaaS (Commands as a Service). Make run commands remotes as local.


## License
GoZoox is released under the [MIT License](./LICENSE).
