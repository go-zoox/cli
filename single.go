package cli

import (
	"fmt"
	"log"
	"os"

	ucli "github.com/urfave/cli/v2"
)

// SingleProgram is a program that has a single command.
type SingleProgram struct {
	cfg    *SingleProgramConfig
	action Action
}

// SingleProgramConfig is the configuration for a SingleProgram.
type SingleProgramConfig struct {
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
	// List of flags to parse
	Flags []Flag
}

// NewSingleProgram creates a new SingleProgram.
func NewSingleProgram(cfg *SingleProgramConfig) *SingleProgram {
	return &SingleProgram{
		cfg: cfg,
	}
}

func (c *SingleProgram) create() (*ucli.App, error) {
	if c.action == nil {
		return nil, fmt.Errorf("command is not set")
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
		Flags:  c.cfg.Flags,
		Action: c.action,
	}, nil
}

// Command sets the action of the program.
func (c *SingleProgram) Command(command Action) {
	c.action = command
}

// Run runs the program.
func (c *SingleProgram) Run(arguments ...[]string) error {
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

// RunWithLog runs the program with log.
func (c *SingleProgram) RunWithLog() {
	if err := c.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
