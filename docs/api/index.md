# API Reference

Complete API reference for the CLI framework.

## Packages

- [SingleProgram](/api/single-program) - Single command CLI applications
- [MultipleProgram](/api/multiple-program) - Multi-command CLI applications
- [Interactive](/api/interactive) - Interactive components API
- [Loading](/api/loading) - Loading components API

## Types

### Context

The `Context` type provides access to command-line flags and arguments.

```go
type Context struct {
	// ... fields
}
```

### Command

The `Command` type represents a CLI command.

```go
type Command struct {
	Name   string
	Usage  string
	Flags  []Flag
	Action Action
}
```

### Flag Types

- `StringFlag` - String flag
- `IntFlag` - Integer flag
- `BoolFlag` - Boolean flag
- `StringSliceFlag` - String slice flag
- `IntSliceFlag` - Integer slice flag

## Functions

### NewSingleProgram

Creates a new single command program.

```go
func NewSingleProgram(cfg *SingleProgramConfig) *SingleProgram
```

### NewMultipleProgram

Creates a new multiple command program.

```go
func NewMultipleProgram(cfg *MultipleProgramConfig) *MultipleProgram
```
