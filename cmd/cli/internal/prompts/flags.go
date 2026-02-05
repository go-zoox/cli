package prompts

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

// FlagInfo contains flag information
type FlagInfo struct {
	Name     string
	Type     string // "string", "int", "bool", "stringSlice"
	Usage    string
	Default  string
	Aliases  []string
	EnvVars  []string
	Required bool
}

// CollectFlags collects flags interactively
func CollectFlags() ([]FlagInfo, error) {
	var flags []FlagInfo

	addMore := true
	for addMore {
		flag, err := collectSingleFlag("")
		if err != nil {
			return nil, err
		}
		flags = append(flags, *flag)

		more, err := interactive.Confirm(
			"Add another flag?",
			&interactive.ConfirmOptions{Default: false},
		)
		if err != nil {
			return nil, err
		}
		addMore = more
	}

	return flags, nil
}

// CollectSingleFlag collects a single flag interactively
func CollectSingleFlag() (*FlagInfo, error) {
	return CollectSingleFlagWithName("")
}

// CollectSingleFlagWithName collects a single flag interactively, using the provided name
func CollectSingleFlagWithName(defaultName string) (*FlagInfo, error) {
	return collectSingleFlag(defaultName)
}

func collectSingleFlag(defaultName string) (*FlagInfo, error) {
	var name string
	var err error

	if defaultName != "" {
		name = defaultName
		fmt.Printf("Flag name: %s\n", name)
	} else {
		name, err = interactive.Text("Flag name:", &interactive.TextOptions{
			Required: true,
		})
		if err != nil {
			return nil, err
		}
	}

	flagType, err := interactive.Select(
		"Flag type:",
		[]interactive.SelectOption{
			{Label: "String", Value: "string"},
			{Label: "Int", Value: "int"},
			{Label: "Bool", Value: "bool"},
			{Label: "String Slice", Value: "stringSlice"},
		},
		&interactive.SelectOptions{Default: "string"},
	)
	if err != nil {
		return nil, err
	}

	usage, err := interactive.Text("Flag usage/description:", &interactive.TextOptions{
		Required: true,
	})
	if err != nil {
		return nil, err
	}

	var defaultValue string
	if flagType != "bool" {
		defaultValue, err = interactive.Text("Default value (optional, press Enter to skip):", nil)
		if err != nil {
			return nil, err
		}
	}

	// Ask for aliases
	hasAlias, err := interactive.Confirm("Add short alias (e.g., -s)?", &interactive.ConfirmOptions{
		Default: false,
	})
	if err != nil {
		return nil, err
	}

	var aliases []string
	if hasAlias {
		alias, err := interactive.Text("Alias (single character, e.g., 's'):", &interactive.TextOptions{
			Required: true,
		})
		if err != nil {
			return nil, err
		}
		aliases = []string{alias}
	}

	// Ask for environment variables
	hasEnvVar, err := interactive.Confirm("Add environment variable support?", &interactive.ConfirmOptions{
		Default: false,
	})
	if err != nil {
		return nil, err
	}

	var envVars []string
	if hasEnvVar {
		envVar, err := interactive.Text("Environment variable name (e.g., MY_FLAG):", &interactive.TextOptions{
			Required: true,
		})
		if err != nil {
			return nil, err
		}
		envVars = []string{envVar}
	}

	return &FlagInfo{
		Name:    name,
		Type:    flagType,
		Usage:   usage,
		Default: defaultValue,
		Aliases: aliases,
		EnvVars: envVars,
	}, nil
}
