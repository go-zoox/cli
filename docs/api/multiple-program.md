# MultipleProgram API

## NewMultipleProgram

Creates a new multiple command program.

```go
func NewMultipleProgram(cfg *MultipleProgramConfig) *MultipleProgram
```

### MultipleProgramConfig

```go
type MultipleProgramConfig struct {
	Name  string
	Usage string
}
```

### MultipleProgram Methods

#### Register

Registers a new command.

```go
func (c *MultipleProgram) Register(name string, cmd *Command)
```

#### Run

Runs the program.

```go
func (c *MultipleProgram) Run()
```

## Example

```go
app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
	Name:  "myapp",
	Usage: "My application",
})

app.Register("hello", &cli.Command{
	Name:  "hello",
	Usage: "Say hello",
	Action: func(ctx *cli.Context) error {
		fmt.Println("Hello!")
		return nil
	},
})

app.Run()
```
