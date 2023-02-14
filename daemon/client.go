package daemon

import (
	"fmt"
	"os"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/core-utils/regexp"
	"github.com/go-zoox/fetch"
)

type Client interface {
	Run() error
}

type client struct {
	app     *fetch.Fetch
	cli     *cli.MultipleProgram
	address string
}

func NewClient(address string) (Client, error) {
	cfg := &fetch.Config{
		BaseURL: address,
	}
	if regexp.Match("^unix://", address) {
		cfg.BaseURL = "http://127.0.0.1"
		cfg.UnixDomainSocket = address
	}

	return &client{
		address: address,
		cli: cli.NewMultipleProgram(&cli.MultipleProgramConfig{
			Name:        "daemon",
			Description: "asdasd",
		}),
		app: fetch.New(cfg),
	}, nil
}

func (c *client) Run() error {
	commands, err := c.getCommands()
	if err != nil {
		return err
	}

	for _, cmd := range commands {
		flags := []cli.Flag{}
		for _, flag := range cmd.Flags {
			switch flag.Type {
			case "string":
				flags = append(flags, &cli.StringFlag{
					Name:    flag.Name,
					Usage:   flag.Usage,
					Aliases: flag.Aliases,
					EnvVars: flag.EnvVars,
					// Value:    flag.Value.(string),
					Required: flag.Required,
				})
			case "bool":
				flags = append(flags, &cli.BoolFlag{
					Name:    flag.Name,
					Usage:   flag.Usage,
					Aliases: flag.Aliases,
					EnvVars: flag.EnvVars,
					// Value:    flag.Value.(bool),
					Required: flag.Required,
				})
			case "int":
				flags = append(flags, &cli.Int64Flag{
					Name:    flag.Name,
					Usage:   flag.Usage,
					Aliases: flag.Aliases,
					EnvVars: flag.EnvVars,
					// Value:    flag.Value.(int64),
					Required: flag.Required,
				})
			default:
				return fmt.Errorf("invalid flag type: %s", flag.Type)
			}
		}

		cmdName := cmd.Name
		c.cli.Register(cmdName, &cli.Command{
			Name:  cmdName,
			Usage: cmd.Usage,
			Flags: flags,
			Action: func(ctx *cli.Context) error {
				output, err := c.execute(cmdName, os.Args[2:])
				if err != nil {
					return err
				}

				fmt.Printf("%s", output)
				return nil
			},
		})
	}

	return c.cli.Run()
}

func (c *client) execute(name string, args []string) (string, error) {
	response, err := c.app.
		Post("/commands/exec", &fetch.Config{
			Body: map[string]any{
				"name": name,
				"args": args,
			},
		}).
		Execute()
	if err != nil {
		return "", err
	}
	if !response.Ok() {
		return "", fmt.Errorf("%s", response.Get("message"))
	}

	return response.Get("result").String(), nil
}

func (c *client) getCommands() (commands []*Command, err error) {
	response, err := c.app.
		Get("/commands").
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get commands: %v", err)
	}

	if !response.Ok() {
		return nil, fmt.Errorf("failed to get commands: [%d] %s", response.Get("code").Int(), response.Get("message").String())
	}

	type DTO struct {
		Commands []*Command `json:"commands"`
	}
	var dto DTO
	if err := response.UnmarshalJSON(&dto); err != nil {
		return nil, fmt.Errorf("failed to unmarshal commands: %v", err)
	}

	return dto.Commands, nil
}
