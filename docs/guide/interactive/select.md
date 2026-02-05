# Select

The `Select` component allows users to choose from a list of options.

## Basic Usage

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	city, err := interactive.Select(
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

	fmt.Println("Your favorite city:", city)
}
```

## Options

### SelectOption

| Field | Type | Description |
|-------|------|-------------|
| `Label` | `string` | The display text for the option |
| `Value` | `string` | The value returned when this option is selected |

### SelectOptions

| Field | Type | Description |
|-------|------|-------------|
| `Default` | `string` | Default selected value |

## Examples

### Simple Select

```go
color, err := interactive.Select(
	"Choose a color:",
	[]interactive.SelectOption{
		{Label: "Red", Value: "red"},
		{Label: "Green", Value: "green"},
		{Label: "Blue", Value: "blue"},
	},
	nil,
)
```

### With Default Value

```go
language, err := interactive.Select(
	"Choose a language:",
	[]interactive.SelectOption{
		{Label: "Go", Value: "go"},
		{Label: "JavaScript", Value: "js"},
		{Label: "Python", Value: "python"},
	},
	&interactive.SelectOptions{
		Default: "go",
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
	// Select framework
	framework, err := interactive.Select(
		"Which framework do you want to use?",
		[]interactive.SelectOption{
			{Label: "React", Value: "react"},
			{Label: "Vue", Value: "vue"},
			{Label: "Angular", Value: "angular"},
			{Label: "Svelte", Value: "svelte"},
		},
		&interactive.SelectOptions{
			Default: "react",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("You selected: %s\n", framework)
}
```
