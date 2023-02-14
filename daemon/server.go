package daemon

import (
	"fmt"
	"net/http"

	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
)

type Server interface {
	Run() error
	Command(cmd *Command, action Action) error
}

type server struct {
	address        string
	app            *zoox.Application
	commands       map[string]*Command
	commandNames   []string
	commandActions map[string]Action
}

func NewServer(address string) (Server, error) {
	return &server{
		address:        address,
		app:            defaults.Application(),
		commands:       make(map[string]*Command),
		commandNames:   make([]string, 0),
		commandActions: make(map[string]Action),
	}, nil
}

func (s *server) Run() error {
	s.app.Post("/commands/exec", func(ctx *zoox.Context) {
		type DTO struct {
			Name string   `json:"name"`
			Args []string `json:"args"`
		}

		var dto DTO
		if err := ctx.BindJSON(&dto); err != nil {
			ctx.Fail(err, http.StatusBadRequest, "invalid data")
			return
		}

		command := s.commands[dto.Name]
		if command == nil {
			ctx.Fail(nil, http.StatusNotFound, "invalid command")
			return
		}

		action := s.commandActions[dto.Name]

		response, err := action(dto.Args...)
		if err != nil {
			ctx.Fail(err, http.StatusInternalServerError, fmt.Sprintf("failed to execute command: %s", err.Error()))
			return
		}

		ctx.Success(response)
	})

	s.app.Get("/commands", func(ctx *zoox.Context) {
		commands := []*Command{}
		for _, name := range s.commandNames {
			commands = append(commands, s.commands[name])
		}

		ctx.JSON(http.StatusOK, zoox.H{
			"commands": commands,
		})
	})

	return s.app.Run(s.address)
}

func (s *server) Command(cmd *Command, action Action) error {
	if _, found := s.commands[cmd.Name]; found {
		return fmt.Errorf("command %s already exists", cmd.Name)
	}

	s.commands[cmd.Name] = cmd
	s.commandActions[cmd.Name] = action
	s.commandNames = append(s.commandNames, cmd.Name)
	return nil
}
