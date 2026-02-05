# SingleProgram API

## NewSingleProgram

Creates a new single command program.

```go
func NewSingleProgram(cfg *SingleProgramConfig) *SingleProgram
```

### SingleProgramConfig

```go
type SingleProgramConfig struct {
	Name                 string
	HelpName             string
	Usage                string
	UsageText            string
	ArgsUsage            string
	Version              string
	Description          string
	EnableBashCompletion bool
	HideHelp             bool
	HideHelpCommand      bool
	HideVersion          bool
	Flags                []Flag
}
```

### SingleProgram Methods

#### Command

Sets the action function for the program.

```go
func (c *SingleProgram) Command(command Action)
```

#### Run

Runs the program.

```go
func (c *SingleProgram) Run()
```

## Example

```go
app := cli.NewSingleProgram(&cli.SingleProgramConfig{
	Name:    "myapp",
	Usage:   "My application",
	Version: "1.0.0",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Usage: "Your name",
		},
	},
})

app.Command(func(ctx *cli.Context) error {
	fmt.Println("Hello,", ctx.String("name"))
	return nil
})

app.Run()
```
