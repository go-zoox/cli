# Password

The `Password` component allows users to enter passwords securely (input is hidden).

## Basic Usage

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	password, err := interactive.Password(
		"Please type your password:",
		&interactive.PasswordOptions{
			Required: true,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Password entered successfully")
}
```

## Options

### PasswordOptions

| Field | Type | Description |
|-------|------|-------------|
| `Required` | `bool` | Whether the password is required (default: `false`) |

## Examples

### Required Password

```go
password, err := interactive.Password(
	"Enter your password:",
	&interactive.PasswordOptions{
		Required: true,
	},
)
```

### Optional Password

```go
password, err := interactive.Password(
	"Enter your password (optional):",
	nil,
)
```

### Password Confirmation

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	password, err := interactive.Password(
		"Enter your password:",
		&interactive.PasswordOptions{
			Required: true,
		},
	)
	if err != nil {
		panic(err)
	}

	confirmPassword, err := interactive.Password(
		"Confirm your password:",
		&interactive.PasswordOptions{
			Required: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if password != confirmPassword {
		fmt.Println("Passwords do not match!")
		return
	}

	fmt.Println("Password set successfully!")
}
```
