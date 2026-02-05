# Multiselect

The `Multiselect` component allows users to select multiple options from a list.

## Basic Usage

```go
package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	items, err := interactive.Multiselect(
		"Select your favorite languages:",
		[]interactive.SelectOption{
			{Label: "Go", Value: "go"},
			{Label: "JavaScript", Value: "js"},
			{Label: "Python", Value: "python"},
			{Label: "Rust", Value: "rust"},
		},
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Selected languages:", items)
}
```

## Options

### SelectOption

| Field | Type | Description |
|-------|------|-------------|
| `Label` | `string` | The display text for the option |
| `Value` | `string` | The value returned when this option is selected |

### MultiselectOptions

| Field | Type | Description |
|-------|------|-------------|
| `Default` | `[]string` | Default selected values |

## Examples

### Simple Multiselect

```go
tags, err := interactive.Multiselect(
	"Select tags:",
	[]interactive.SelectOption{
		{Label: "Frontend", Value: "frontend"},
		{Label: "Backend", Value: "backend"},
		{Label: "DevOps", Value: "devops"},
	},
	nil,
)
```

### With Default Values

```go
frameworks, err := interactive.Multiselect(
	"Select frameworks:",
	[]interactive.SelectOption{
		{Label: "React", Value: "react"},
		{Label: "Vue", Value: "vue"},
		{Label: "Angular", Value: "angular"},
	},
	&interactive.MultiselectOptions{
		Default: []string{"react", "vue"},
	},
)
```

### Complete Example

```go
package main

import (
	"fmt"
	"strings"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	features, err := interactive.Multiselect(
		"Select features to enable:",
		[]interactive.SelectOption{
			{Label: "Authentication", Value: "auth"},
			{Label: "Database", Value: "db"},
			{Label: "Caching", Value: "cache"},
			{Label: "Logging", Value: "logging"},
			{Label: "Monitoring", Value: "monitoring"},
		},
		&interactive.MultiselectOptions{
			Default: []string{"auth", "db"},
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Selected features: %s\n", strings.Join(features, ", "))
}
```
