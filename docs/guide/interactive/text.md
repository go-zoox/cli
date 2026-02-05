# Text Input

The `Text` component allows you to prompt users for text input in an interactive way.

## Basic Usage

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	name, err := interactive.Text("What is your name?", &interactive.TextOptions{
		Default:  "Zero",
		Required: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Your name is:", name)
}
```

## Options

### TextOptions

| Field | Type | Description |
|-------|------|-------------|
| `Default` | `string` | Default value if user presses Enter without typing |
| `Required` | `bool` | Whether the input is required (default: `false`) |
| `Placeholder` | `string` | Placeholder text shown in the input field |

## Examples

### With Default Value

```go
name, err := interactive.Text("What is your name?", &interactive.TextOptions{
	Default: "Anonymous",
})
```

### Required Input

```go
email, err := interactive.Text("What is your email?", &interactive.TextOptions{
	Required: true,
})
```

### With Placeholder

```go
username, err := interactive.Text("Enter username:", &interactive.TextOptions{
	Placeholder: "username",
	Required:    true,
})
```

### Complete Example

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	// Required input with default
	name, err := interactive.Text("What is your name?", &interactive.TextOptions{
		Default:  "Guest",
		Required: true,
	})
	if err != nil {
		panic(err)
	}

	// Optional input
	email, err := interactive.Text("What is your email? (optional)", &interactive.TextOptions{
		Placeholder: "user@example.com",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s\n", name)
	if email != "" {
		fmt.Printf("Email: %s\n", email)
	}
}
```
