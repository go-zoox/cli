package cli

import (
	"fmt"
	"log"
	"os"

	ucli "github.com/urfave/cli/v2"
)

// MultipleProgram is a program that has multiple commands.
type MultipleProgram struct {
	cfg      *MultipleProgramConfig
	commands map[string]*Command
}

// MultipleProgramConfig is the configuration for a MultipleProgram.
type MultipleProgramConfig struct {
	// The name of the program. Defaults to path.Base(os.Args[0])
	Name string
	// Full name of command for help, defaults to Name
	HelpName string
	// Description of the program.
	Usage string
	// Text to override the USAGE section of help
	UsageText string
	// Description of the program argument format.
	ArgsUsage string
	// Version of the program
	Version string
	// Description of the program
	Description string
	// Boolean to enable bash completion commands
	EnableBashCompletion bool
	// Boolean to hide built-in help command and help flag
	HideHelp bool
	// Boolean to hide built-in help command but keep help flag.
	// Ignored if HideHelp is true.
	HideHelpCommand bool
	// Boolean to hide built-in version flag and the VERSION section of help
	HideVersion bool
}

// NewMultipleProgram creates a new MultipleProgram.
func NewMultipleProgram(cfg *MultipleProgramConfig) *MultipleProgram {
	return &MultipleProgram{
		cfg:      cfg,
		commands: make(map[string]*Command),
	}
}

func (c *MultipleProgram) create() (*ucli.App, error) {
	if len(c.commands) == 0 {
		return nil, fmt.Errorf("command is not set")
	}

	commands := make([]*ucli.Command, 0, len(c.commands))
	for _, cmd := range c.commands {
		commands = append(commands, cmd)
	}

	return &ucli.App{
		Name:                 c.cfg.Name,
		Usage:                c.cfg.Usage,
		UsageText:            c.cfg.UsageText,
		ArgsUsage:            c.cfg.ArgsUsage,
		Version:              c.cfg.Version,
		Description:          c.cfg.Description,
		EnableBashCompletion: c.cfg.EnableBashCompletion,
		HideHelp:             c.cfg.HideHelp,
		HideHelpCommand:      c.cfg.HideHelpCommand,
		HideVersion:          c.cfg.HideVersion,
		//
		Commands: commands,
	}, nil
}

// Register registers a command.
func (c *MultipleProgram) Register(name string, cmd *Command) error {
	if _, ok := c.commands[name]; ok {
		return fmt.Errorf("command %s is already registered", name)
	}

	c.commands[name] = cmd
	return nil
}

// Run runs the program.
func (c *MultipleProgram) Run(arguments ...[]string) error {
	argumentsX := os.Args
	if len(arguments) > 0 && len(arguments[0]) > 0 {
		argumentsX = arguments[0]
	}

	app, err := c.create()
	if err != nil {
		return err
	}

	return app.Run(argumentsX)
}

// Serve runs the program with log.
func (c *MultipleProgram) Serve() {
	if err := c.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
