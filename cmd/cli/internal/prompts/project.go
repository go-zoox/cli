package prompts

import (
	"fmt"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/cli/interactive"
)

// ProjectInfo contains project information
type ProjectInfo struct {
	Name        string
	Usage       string
	Version     string
	Type        string // "single" or "multiple"
	Description string
}

// CollectProjectInfo collects project information interactively
func CollectProjectInfo(ctx *cli.Context) (*ProjectInfo, error) {
	return CollectProjectInfoWithDefaults(ctx, "", "", "", "")
}

// CollectProjectInfoWithDefaults collects project information, using provided defaults and only asking for missing values
func CollectProjectInfoWithDefaults(ctx *cli.Context, defaultName, defaultUsage, defaultVersion, defaultType string) (*ProjectInfo, error) {
	var name, usage, version, cliType string
	var err error

	// Name
	if defaultName != "" {
		name = defaultName
		fmt.Printf("Project name: %s\n", name)
	} else {
		name, err = interactive.Text("Project name:", &interactive.TextOptions{
			Required: true,
		})
		if err != nil {
			return nil, err
		}
	}

	// Usage
	if defaultUsage != "" {
		usage = defaultUsage
		fmt.Printf("Project usage/description: %s\n", usage)
	} else {
		usage, err = interactive.Text("Project usage/description:", &interactive.TextOptions{
			Required: true,
		})
		if err != nil {
			return nil, err
		}
	}

	// Version
	if defaultVersion != "" {
		version = defaultVersion
		fmt.Printf("Version: %s\n", version)
	} else {
		version, err = interactive.Text("Version:", &interactive.TextOptions{
			Default: "0.1.0",
		})
		if err != nil {
			return nil, err
		}
	}

	// Type
	if defaultType != "" {
		if defaultType != "single" && defaultType != "multiple" {
			return nil, fmt.Errorf("invalid type: %s (must be 'single' or 'multiple')", defaultType)
		}
		cliType = defaultType
		fmt.Printf("CLI type: %s\n", cliType)
	} else {
		cliType, err = interactive.Select(
			"CLI type:",
			[]interactive.SelectOption{
				{Label: "Single Command", Value: "single"},
				{Label: "Multiple Commands", Value: "multiple"},
			},
			&interactive.SelectOptions{
				Default: "single",
			},
		)
		if err != nil {
			return nil, err
		}
	}

	return &ProjectInfo{
		Name:    name,
		Usage:   usage,
		Version: version,
		Type:    cliType,
	}, nil
}

// Confirm prompts for confirmation
func Confirm(message string, defaultValue bool) (bool, error) {
	return interactive.Confirm(message, &interactive.ConfirmOptions{
		Default: defaultValue,
	})
}
