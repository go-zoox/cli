# Confirm

The `Confirm` component allows users to answer yes/no questions.

## Basic Usage

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	like, err := interactive.Confirm(
		"Do you like the book?",
		&interactive.ConfirmOptions{
			Default: false,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your answer:", like)
}
```

## Options

### ConfirmOptions

| Field | Type | Description |
|-------|------|-------------|
| `Default` | `bool` | Default value if user presses Enter without typing (default: `false`) |

## Examples

### Simple Confirm

```go
proceed, err := interactive.Confirm("Do you want to continue?", nil)
if err != nil {
	panic(err)
}

if proceed {
	fmt.Println("Continuing...")
} else {
	fmt.Println("Cancelled.")
}
```

### With Default Value

```go
// Default to yes
delete, err := interactive.Confirm(
	"Are you sure you want to delete this file?",
	&interactive.ConfirmOptions{
		Default: true,
	},
)
```

### Complete Example

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	// Confirm deletion
	delete, err := interactive.Confirm(
		"Are you sure you want to delete this file?",
		&interactive.ConfirmOptions{
			Default: false,
		},
	)
	if err != nil {
		panic(err)
	}

	if delete {
		fmt.Println("File deleted.")
	} else {
		fmt.Println("Deletion cancelled.")
	}

	// Confirm installation
	install, err := interactive.Confirm(
		"Do you want to install dependencies?",
		&interactive.ConfirmOptions{
			Default: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if install {
		fmt.Println("Installing dependencies...")
	}
}
```
