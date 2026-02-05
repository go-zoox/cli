# Interactive Examples

This page contains examples of using interactive components.

## Text Input

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

## Select

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

## Confirm

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

## Password

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
