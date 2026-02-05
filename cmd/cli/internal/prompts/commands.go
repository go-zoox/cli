package prompts

import (
	"github.com/go-zoox/cli/interactive"
)

// CommandInfo contains command information
type CommandInfo struct {
	Name        string
	Usage       string
	Flags       []FlagInfo
	Description string
}

// CollectCommands collects commands interactively
func CollectCommands() ([]CommandInfo, error) {
	var commands []CommandInfo

	addMore := true
	for addMore {
		cmd, err := collectSingleCommand()
		if err != nil {
			return nil, err
		}
		commands = append(commands, *cmd)

		more, err := interactive.Confirm(
			"Add another command?",
			&interactive.ConfirmOptions{Default: false},
		)
		if err != nil {
			return nil, err
		}
		addMore = more
	}

	return commands, nil
}

// CollectSingleCommand collects a single command interactively
func CollectSingleCommand() (*CommandInfo, error) {
	return collectSingleCommand()
}

func collectSingleCommand() (*CommandInfo, error) {
	name, err := interactive.Text("Command name:", &interactive.TextOptions{
		Required: true,
	})
	if err != nil {
		return nil, err
	}

	usage, err := interactive.Text("Command usage/description:", &interactive.TextOptions{
		Required: true,
	})
	if err != nil {
		return nil, err
	}

	// Ask if command needs flags
	hasFlags, err := interactive.Confirm("Add flags to this command?", &interactive.ConfirmOptions{
		Default: false,
	})
	if err != nil {
		return nil, err
	}

	var flags []FlagInfo
	if hasFlags {
		flags, err = CollectFlags()
		if err != nil {
			return nil, err
		}
	}

	return &CommandInfo{
		Name:  name,
		Usage: usage,
		Flags: flags,
	}, nil
}
