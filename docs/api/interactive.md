# Interactive API

Interactive components for user input.

## Text

Prompt for text input.

```go
func Text(prompt string, opts *TextOptions) (string, error)
```

### TextOptions

```go
type TextOptions struct {
	Default     string
	Required    bool
	Placeholder string
}
```

## Select

Single selection from options.

```go
func Select(prompt string, options []SelectOption, opts *SelectOptions) (string, error)
```

### SelectOption

```go
type SelectOption struct {
	Label string
	Value string
}
```

### SelectOptions

```go
type SelectOptions struct {
	Default string
}
```

## Confirm

Yes/No confirmation.

```go
func Confirm(prompt string, opts *ConfirmOptions) (bool, error)
```

### ConfirmOptions

```go
type ConfirmOptions struct {
	Default bool
}
```

## Password

Secure password input.

```go
func Password(prompt string, opts *PasswordOptions) (string, error)
```

### PasswordOptions

```go
type PasswordOptions struct {
	Required bool
}
```

## Multiselect

Multiple selection from options.

```go
func Multiselect(prompt string, options []SelectOption, opts *MultiselectOptions) ([]string, error)
```

### MultiselectOptions

```go
type MultiselectOptions struct {
	Default []string
}
```
