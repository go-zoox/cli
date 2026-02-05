package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/cli/cmd/cli/internal/adder"
	"github.com/go-zoox/cli/cmd/cli/internal/config"
	"github.com/go-zoox/cli/cmd/cli/internal/devtools"
	"github.com/go-zoox/cli/cmd/cli/internal/docs"
	"github.com/go-zoox/cli/cmd/cli/internal/formatter"
	"github.com/go-zoox/cli/cmd/cli/internal/generator"
	"github.com/go-zoox/cli/cmd/cli/internal/linter"
	"github.com/go-zoox/cli/cmd/cli/internal/listing"
	"github.com/go-zoox/cli/cmd/cli/internal/migrate"
	"github.com/go-zoox/cli/cmd/cli/internal/prompts"
	"github.com/go-zoox/cli/cmd/cli/internal/remover"
	"github.com/go-zoox/cli/cmd/cli/internal/template"
	"github.com/go-zoox/cli/cmd/cli/internal/testgen"
	"github.com/go-zoox/cli/cmd/cli/internal/update"
	"github.com/go-zoox/cli/cmd/cli/internal/upgrade"
	"github.com/go-zoox/cli/cmd/cli/internal/validate"
	"github.com/go-zoox/cli/interactive"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:  "cli",
		Usage: "CLI framework tools and utilities",
	})

	app.Register(&cli.Command{
		Name:  "init",
		Usage: "Initialize a new CLI project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Usage:   "Project name",
				Aliases: []string{"n"},
			},
			&cli.StringFlag{
				Name:    "type",
				Usage:   "CLI type (single|multiple)",
				Aliases: []string{"t"},
			},
			&cli.StringFlag{
				Name:    "output",
				Usage:   "Output directory",
				Aliases: []string{"o"},
				Value:   ".",
			},
			&cli.BoolFlag{
				Name:    "skip-interactive",
				Usage:   "Skip interactive prompts (use flags only)",
				Aliases: []string{"s"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runInit(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "add",
		Usage: "Add a command or flag to an existing CLI project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "command",
				Usage:   "Command name to add (for multiple commands CLI)",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:    "flag",
				Usage:   "Flag name to add",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:  "to-command",
				Usage: "Add flag to specific command (for multiple commands CLI)",
			},
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runAdd(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "list",
		Usage: "List commands and flags of an existing CLI project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.StringFlag{
				Name:    "format",
				Usage:   "Output format (table|json|yaml)",
				Aliases: []string{"f"},
				Value:   "table",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runList(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "remove",
		Usage: "Remove a command or flag from an existing CLI project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "command",
				Usage:   "Command name to remove (for multiple commands CLI)",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:    "flag",
				Usage:   "Flag name to remove",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:  "from-command",
				Usage: "Remove flag from specific command (for multiple commands CLI)",
			},
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runRemove(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "validate",
		Usage: "Validate an existing CLI project (structure and basic rules)",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.BoolFlag{
				Name:  "fix",
				Usage: "Attempt to fix issues (gofmt main.go)",
				Value: false,
			},
			&cli.StringFlag{
				Name:    "format",
				Usage:   "Output format (text|json)",
				Aliases: []string{"f"},
				Value:   "text",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runValidate(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "template",
		Usage: "Manage CLI project templates",
		Flags: []cli.Flag{},
		Subcommands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all available templates",
				Action: func(ctx *cli.Context) error {
					return runTemplateList(ctx)
				},
			},
			{
				Name:  "add",
				Usage: "Add a new template",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Template name",
						Aliases:  []string{"n"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Template path (local directory or github:user/repo)",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    "type",
						Usage:   "Template type (local|github|gitlab)",
						Aliases: []string{"t"},
						Value:   "local",
					},
				},
				Action: func(ctx *cli.Context) error {
					return runTemplateAdd(ctx)
				},
			},
			{
				Name:  "remove",
				Usage: "Remove a template",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Template name to remove",
						Aliases:  []string{"n"},
						Required: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					return runTemplateRemove(ctx)
				},
			},
			{
				Name:  "use",
				Usage: "Use a template to initialize a project",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Usage:    "Template name to use",
						Aliases:  []string{"n"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    "output",
						Usage:   "Output directory",
						Aliases: []string{"o"},
						Value:   ".",
					},
				},
				Action: func(ctx *cli.Context) error {
					return runTemplateUse(ctx)
				},
			},
		},
	})

	app.Register(&cli.Command{
		Name:  "format",
		Usage: "Format Go code using gofmt and goimports",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.BoolFlag{
				Name:    "check",
				Usage:   "Check if code is formatted (don't modify)",
				Aliases: []string{"c"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runFormat(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "lint",
		Usage: "Lint Go code using golangci-lint or go vet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.BoolFlag{
				Name:    "fix",
				Usage:   "Automatically fix issues",
				Aliases: []string{"f"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runLint(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "test",
		Usage: "Generate test files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.StringFlag{
				Name:    "command",
				Usage:   "Generate test for specific command",
				Aliases: []string{"c"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runTest(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "upgrade",
		Usage: "Upgrade CLI framework version",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.StringFlag{
				Name:    "to",
				Usage:   "Upgrade to specific version",
				Aliases: []string{"t"},
			},
			&cli.BoolFlag{
				Name:    "check",
				Usage:   "Check for available upgrades",
				Aliases: []string{"c"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runUpgrade(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "migrate",
		Usage: "Migrate project structure",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.StringFlag{
				Name:     "to",
				Usage:    "Target structure (multiple)",
				Aliases:  []string{"t"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			return runMigrate(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "config",
		Usage: "Manage CLI tool configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "key",
				Usage:   "Config key",
				Aliases: []string{"k"},
			},
			&cli.StringFlag{
				Name:    "value",
				Usage:   "Config value (for set)",
				Aliases: []string{"v"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runConfig(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "build",
		Usage: "Build CLI project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.StringFlag{
				Name:    "output",
				Usage:   "Output binary path",
				Aliases: []string{"o"},
			},
			&cli.StringFlag{
				Name:    "platform",
				Usage:   "Target platform (e.g., linux/amd64)",
				Aliases: []string{"p"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runBuild(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "run",
		Usage: "Run CLI project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runRun(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "watch",
		Usage: "Watch for file changes and rebuild",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.BoolFlag{
				Name:    "build",
				Usage:   "Rebuild on changes",
				Aliases: []string{"b"},
			},
			&cli.BoolFlag{
				Name:    "test",
				Usage:   "Run tests on changes",
				Aliases: []string{"t"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return runWatch(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "docs",
		Usage: "Generate documentation",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.StringFlag{
				Name:    "output",
				Usage:   "Output directory",
				Aliases: []string{"o"},
			},
			&cli.StringFlag{
				Name:    "format",
				Usage:   "Output format (markdown)",
				Aliases: []string{"f"},
				Value:   "markdown",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runDocs(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "man",
		Usage: "Generate man page",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runMan(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "completion",
		Usage: "Generate shell completion script",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "shell",
				Usage:    "Shell type (bash|zsh|fish|powershell)",
				Aliases:  []string{"s"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runCompletion(ctx)
		},
	})

	app.Register(&cli.Command{
		Name:  "version",
		Usage: "Show version information",
		Action: func(ctx *cli.Context) error {
			fmt.Println("cli tool version 0.1.0")
			return nil
		},
	})

	app.Register(&cli.Command{
		Name:  "update",
		Usage: "Update command or flag properties in an existing CLI project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Usage:   "Project directory (default: current directory)",
				Aliases: []string{"d"},
				Value:   ".",
			},
			&cli.StringFlag{
				Name:    "command",
				Usage:   "Command name to update",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:  "usage",
				Usage: "New usage/description for the command",
			},
			&cli.StringFlag{
				Name:    "flag",
				Usage:   "Flag name to update",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:  "flag-usage",
				Usage: "New usage/description for the flag",
			},
			&cli.StringFlag{
				Name:  "flag-default",
				Usage: "New default value for the flag",
			},
			&cli.StringFlag{
				Name:  "from-command",
				Usage: "When updating a flag, limit to this command (for multiple commands CLI)",
			},
		},
		Action: func(ctx *cli.Context) error {
			return runUpdate(ctx)
		},
	})

	app.Run()
}

func runInit(ctx *cli.Context) error {
	// 显示欢迎信息
	fmt.Println("🚀 CLI Tool - Generate CLI applications quickly!")
	fmt.Println()

	skipInteractive := ctx.Bool("skip-interactive")

	// 获取命令行参数
	name := ctx.String("name")
	cliType := ctx.String("type")
	outputDir := ctx.String("output")

	// 使用默认值或通过交互式询问补充缺失的参数
	var info *prompts.ProjectInfo
	var err error

	if skipInteractive {
		// 完全非交互式模式 - 所有必需参数必须提供
		if name == "" {
			return fmt.Errorf("--name is required when using --skip-interactive")
		}
		if cliType == "" {
			cliType = "single"
		}
		if cliType != "single" && cliType != "multiple" {
			return fmt.Errorf("--type must be 'single' or 'multiple'")
		}

		info = &prompts.ProjectInfo{
			Name:    name,
			Usage:   fmt.Sprintf("%s is a CLI application", name),
			Version: "0.1.0",
			Type:    cliType,
		}
	} else {
		// 混合模式：使用提供的参数，询问缺失的参数
		usage := ""
		version := "0.1.0"

		info, err = prompts.CollectProjectInfoWithDefaults(ctx, name, usage, version, cliType)
		if err != nil {
			return err
		}
	}

	// 使用固定的示例 flags 和 commands
	var flags []prompts.FlagInfo
	var commands []prompts.CommandInfo

	if info.Type == "single" {
		// 单命令模式：提供固定的示例 flags
		flags = []prompts.FlagInfo{
			{
				Name:     "name",
				Type:     "string",
				Usage:    "Your name",
				Default:  "World",
				Aliases:  []string{"n"},
				EnvVars:  []string{},
				Required: false,
			},
			{
				Name:     "verbose",
				Type:     "bool",
				Usage:    "Enable verbose output",
				Default:  "",
				Aliases:  []string{"v"},
				EnvVars:  []string{},
				Required: false,
			},
		}
	} else {
		// 多命令模式：提供固定的示例 commands
		commands = []prompts.CommandInfo{
			{
				Name:  "list",
				Usage: "List items",
				Flags: []prompts.FlagInfo{
					{
						Name:     "all",
						Type:     "bool",
						Usage:    "Show all items",
						Default:  "",
						Aliases:  []string{"a"},
						EnvVars:  []string{},
						Required: false,
					},
				},
			},
			{
				Name:  "create",
				Usage: "Create a new item",
				Flags: []prompts.FlagInfo{
					{
						Name:     "name",
						Type:     "string",
						Usage:    "Item name",
						Default:  "",
						Aliases:  []string{"n"},
						EnvVars:  []string{},
						Required: true,
					},
				},
			},
		}
	}

	// 生成代码
	if outputDir == "" || outputDir == "." {
		outputDir = info.Name
	}

	gen := generator.NewGenerator(outputDir)

	if info.Type == "single" {
		if err := gen.GenerateSingleCommand(info, flags); err != nil {
			return fmt.Errorf("failed to generate single command: %w", err)
		}
	} else {
		if err := gen.GenerateMultipleCommands(info, commands); err != nil {
			return fmt.Errorf("failed to generate multiple commands: %w", err)
		}
	}

	// 显示完成信息
	fmt.Printf("\n✅ Project generated successfully in '%s'!\n", outputDir)
	fmt.Println("\nNext steps:")
	fmt.Printf("  cd %s\n", outputDir)
	fmt.Println("  go mod tidy")
	fmt.Println("  go run main.go")

	return nil
}

func runAdd(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Check if main.go exists
	mainPath := filepath.Join(absDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s. Make sure you're in a CLI project directory", absDir)
	}

	commandName := ctx.String("command")
	flagName := ctx.String("flag")
	toCommand := ctx.String("to-command")

	// 确定要添加什么
	var addType string
	if commandName != "" {
		addType = "command"
	} else if flagName != "" {
		addType = "flag"
	} else {
		// 交互式询问要添加什么
		choice, err := prompts.Select(
			"What would you like to add?",
			[]prompts.SelectOption{
				{Label: "Command", Value: "command"},
				{Label: "Flag", Value: "flag"},
			},
			&interactive.SelectOptions{},
		)
		if err != nil {
			return err
		}
		addType = choice
	}

	if addType == "command" {
		// 添加命令
		var cmd *prompts.CommandInfo
		if commandName != "" {
			// 提供了命令名，询问其他信息
			fmt.Printf("Command name: %s\n", commandName)
			usage, err := prompts.Text("Command usage/description:", &interactive.TextOptions{
				Required: true,
			})
			if err != nil {
				return err
			}

			// 询问是否需要 flags
			hasFlags, err := prompts.Confirm("Add flags to this command?", false)
			if err != nil {
				return err
			}

			var flags []prompts.FlagInfo
			if hasFlags {
				flags, err = prompts.CollectFlags()
				if err != nil {
					return err
				}
			}

			cmd = &prompts.CommandInfo{
				Name:  commandName,
				Usage: usage,
				Flags: flags,
			}
		} else {
			// 完全交互式收集命令信息
			cmd, err = prompts.CollectSingleCommand()
			if err != nil {
				return err
			}
		}

		adder := adder.NewAdder(absDir)
		if err := adder.AddCommand(cmd); err != nil {
			return fmt.Errorf("failed to add command: %w", err)
		}

		fmt.Printf("\n✅ Command '%s' added successfully!\n", cmd.Name)
		return nil
	} else {
		// 添加 flag
		var flag *prompts.FlagInfo
		if flagName != "" {
			// 提供了 flag 名，询问其他信息
			fmt.Printf("Flag name: %s\n", flagName)
			flag, err = prompts.CollectSingleFlagWithName(flagName)
			if err != nil {
				return err
			}
		} else {
			// 完全交互式收集 flag 信息
			flag, err = prompts.CollectSingleFlag()
			if err != nil {
				return err
			}
		}

		// 询问添加到哪个命令
		var targetCommand string
		if toCommand != "" {
			targetCommand = toCommand
			fmt.Printf("Add flag to command: %s\n", targetCommand)
		} else {
			targetCommand, err = prompts.Text("Add flag to which command? (leave empty for global flags):", nil)
			if err != nil {
				return err
			}
		}

		adder := adder.NewAdder(absDir)
		if err := adder.AddFlag(flag, targetCommand); err != nil {
			return fmt.Errorf("failed to add flag: %w", err)
		}

		fmt.Printf("\n✅ Flag '%s' added successfully!\n", flag.Name)
		return nil
	}
}

func runList(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	mainPath := filepath.Join(absDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s. Make sure you're in a CLI project directory", absDir)
	}

	info, err := listing.List(absDir)
	if err != nil {
		return err
	}

	output, err := listing.FormatOutput(info, ctx.String("format"))
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}

func runRemove(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	mainPath := filepath.Join(absDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s. Make sure you're in a CLI project directory", absDir)
	}

	commandName := ctx.String("command")
	flagName := ctx.String("flag")
	fromCommand := ctx.String("from-command")

	// Determine what to remove
	var removeType string
	if commandName != "" {
		removeType = "command"
	} else if flagName != "" {
		removeType = "flag"
	} else {
		choice, err := prompts.Select(
			"What would you like to remove?",
			[]prompts.SelectOption{
				{Label: "Command", Value: "command"},
				{Label: "Flag", Value: "flag"},
			},
			&interactive.SelectOptions{},
		)
		if err != nil {
			return err
		}
		removeType = choice
	}

	rm := remover.NewRemover(absDir)

	if removeType == "command" {
		var name string
		if commandName != "" {
			name = commandName
			fmt.Printf("Command to remove: %s\n", name)
		} else {
			name, err = prompts.Text("Command name to remove:", &interactive.TextOptions{Required: true})
			if err != nil {
				return err
			}
		}

		confirm, err := prompts.Confirm(fmt.Sprintf("Remove command '%s'?", name), false)
		if err != nil {
			return err
		}
		if !confirm {
			fmt.Println("Cancelled.")
			return nil
		}

		if err := rm.RemoveCommand(name); err != nil {
			return err
		}

		fmt.Printf("\n✅ Command '%s' removed successfully!\nBackup created at main.go.bak\n", name)
		return nil
	}

	// remove flag
	var flag string
	if flagName != "" {
		flag = flagName
		fmt.Printf("Flag to remove: %s\n", flag)
	} else {
		flag, err = prompts.Text("Flag name to remove:", &interactive.TextOptions{Required: true})
		if err != nil {
			return err
		}
	}

	var targetCmd string
	if fromCommand != "" {
		targetCmd = fromCommand
		fmt.Printf("From command: %s\n", targetCmd)
	} else {
		targetCmd, err = prompts.Text("From which command? (leave empty for global flags):", nil)
		if err != nil {
			return err
		}
	}

	confirm, err := prompts.Confirm(fmt.Sprintf("Remove flag '%s'%s?", flag, func() string {
		if targetCmd != "" {
			return fmt.Sprintf(" from command '%s'", targetCmd)
		}
		return ""
	}()), false)
	if err != nil {
		return err
	}
	if !confirm {
		fmt.Println("Cancelled.")
		return nil
	}

	if err := rm.RemoveFlag(flag, targetCmd); err != nil {
		return err
	}

	fmt.Printf("\n✅ Flag '%s' removed successfully! Backup created at main.go.bak\n", flag)
	return nil
}

func runValidate(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	mainPath := filepath.Join(absDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s. Make sure you're in a CLI project directory", absDir)
	}

	res, err := validate.Validate(absDir, ctx.Bool("fix"))
	if err != nil {
		return err
	}

	format := ctx.String("format")
	if format == "json" {
		type validateOutput struct {
			Issues []string `json:"issues"`
		}
		out, err := json.MarshalIndent(validateOutput{Issues: res.Issues}, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(out))
		return nil
	}

	if len(res.Issues) == 0 {
		fmt.Println("✅ Validation passed. No issues found.")
	} else {
		fmt.Println("⚠️ Validation issues:")
		for _, issue := range res.Issues {
			fmt.Printf("- %s\n", issue)
		}
	}

	if ctx.Bool("fix") {
		fmt.Println("Attempted to format main.go (gofmt).")
	}

	return nil
}

func runUpdate(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	mainPath := filepath.Join(absDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s. Make sure you're in a CLI project directory", absDir)
	}

	commandName := ctx.String("command")
	flagName := ctx.String("flag")
	fromCommand := ctx.String("from-command")
	newUsage := ctx.String("usage")
	flagUsage := ctx.String("flag-usage")
	flagDefault := ctx.String("flag-default")

	if commandName == "" && flagName == "" {
		return fmt.Errorf("please specify --command or --flag")
	}

	up := update.NewUpdater(absDir)

	if commandName != "" && flagName == "" {
		if newUsage == "" {
			return fmt.Errorf("please specify --usage to update command usage")
		}
		if err := up.UpdateCommand(commandName, newUsage); err != nil {
			return err
		}
		fmt.Printf("\n✅ Command '%s' updated successfully!\n", commandName)
		return nil
	}

	if flagName != "" {
		if flagUsage == "" && flagDefault == "" {
			return fmt.Errorf("please specify --flag-usage or --flag-default")
		}
		if err := up.UpdateFlag(flagName, fromCommand, flagUsage, flagDefault); err != nil {
			return err
		}
		target := "global"
		if fromCommand != "" {
			target = fmt.Sprintf("command '%s'", fromCommand)
		}
		fmt.Printf("\n✅ Flag '%s' updated successfully in %s!\n", flagName, target)
		return nil
	}

	return fmt.Errorf("nothing to update")
}

func runUpgrade(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	version := ctx.String("to")
	checkOnly := ctx.Bool("check")

	up := upgrade.NewUpgrader(absDir)

	if checkOnly {
		return up.Check()
	}

	return up.Upgrade(version)
}

func runMigrate(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	target := ctx.String("to")
	if target != "multiple" {
		return fmt.Errorf("only migration to 'multiple' is supported currently")
	}

	migrator := migrate.NewMigrator(absDir)
	return migrator.MigrateToMultiple()
}

func runConfig(ctx *cli.Context) error {
	manager, err := config.NewManager()
	if err != nil {
		return err
	}

	key := ctx.String("key")
	value := ctx.String("value")

	if key != "" && value != "" {
		// Set config
		return manager.Set(key, value)
	}

	if key != "" {
		// Get config
		val, err := manager.Get(key)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", val)
		return nil
	}

	// List all config
	all := manager.List()
	if len(all) == 0 {
		fmt.Println("No configuration found.")
		return nil
	}

	fmt.Println("Configuration:")
	for k, v := range all {
		fmt.Printf("  %s = %v\n", k, v)
	}

	return nil
}

func runBuild(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	outputPath := ctx.String("output")
	platform := ctx.String("platform")

	builder := devtools.NewBuilder(absDir)
	return builder.Build(outputPath, platform)
}

func runRun(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	runner := devtools.NewRunner(absDir)
	// Get remaining args after flags
	args := ctx.Args().Slice()
	return runner.Run(args)
}

func runWatch(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	build := ctx.Bool("build")
	test := ctx.Bool("test")

	watcher := devtools.NewWatcher(absDir)
	return watcher.Watch(build, test)
}

func runDocs(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	outputDir := ctx.String("output")
	format := ctx.String("format")

	gen := docs.NewGenerator(absDir)
	return gen.Generate(outputDir, format)
}

func runMan(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	gen := docs.NewManGenerator(absDir)
	return gen.Generate()
}

func runCompletion(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	shell := ctx.String("shell")

	gen := docs.NewCompletionGenerator(absDir)
	script, err := gen.Generate(shell)
	if err != nil {
		return err
	}

	fmt.Print(script)
	return nil
}

func runVersion(ctx *cli.Context) error {
	fmt.Println("cli tool version 0.1.0")
	fmt.Println("CLI framework: github.com/go-zoox/cli")
	return nil
}

func runTemplateList(ctx *cli.Context) error {
	manager, err := template.NewManager()
	if err != nil {
		return err
	}

	templates, err := manager.List()
	if err != nil {
		return err
	}

	if len(templates) == 0 {
		fmt.Println("No templates found.")
		return nil
	}

	fmt.Println("Available templates:")
	fmt.Println()
	for _, t := range templates {
		fmt.Printf("  %s - %s (%s)\n", t.Name, t.Description, t.Type)
	}

	return nil
}

func runTemplateAdd(ctx *cli.Context) error {
	name := ctx.String("name")
	path := ctx.String("path")
	templateType := ctx.String("type")

	manager, err := template.NewManager()
	if err != nil {
		return err
	}

	if err := manager.Add(name, path, templateType); err != nil {
		return err
	}

	fmt.Printf("✅ Template '%s' added successfully!\n", name)
	return nil
}

func runTemplateRemove(ctx *cli.Context) error {
	name := ctx.String("name")

	manager, err := template.NewManager()
	if err != nil {
		return err
	}

	if err := manager.Remove(name); err != nil {
		return err
	}

	fmt.Printf("✅ Template '%s' removed successfully!\n", name)
	return nil
}

func runTemplateUse(ctx *cli.Context) error {
	name := ctx.String("name")
	outputDir := ctx.String("output")

	loader, err := template.NewLoader()
	if err != nil {
		return err
	}

	_, err = loader.LoadTemplate(name)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	// For now, using a template just means using the default generator
	// In the future, this will apply template files
	fmt.Printf("✅ Using template '%s' to initialize project in '%s'\n", name, outputDir)
	fmt.Println("Note: Template system is in development. Using default generator for now.")

	return nil
}

func runFormat(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	checkOnly := ctx.Bool("check")

	formatter := formatter.NewFormatter(absDir)
	if err := formatter.Format(checkOnly); err != nil {
		if checkOnly {
			return err
		}
		return fmt.Errorf("formatting failed: %w", err)
	}

	if checkOnly {
		fmt.Println("✅ Code is properly formatted")
	} else {
		fmt.Println("✅ Code formatted successfully")
	}

	return nil
}

func runLint(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	fix := ctx.Bool("fix")

	linter := linter.NewLinter(absDir)
	if err := linter.Lint(fix); err != nil {
		return err
	}

	fmt.Println("✅ Linting passed")
	return nil
}

func runTest(ctx *cli.Context) error {
	projectDir := ctx.String("dir")
	absDir, err := filepath.Abs(projectDir)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	commandName := ctx.String("command")

	gen := testgen.NewTestGenerator(absDir)
	if err := gen.Generate(commandName); err != nil {
		return err
	}

	if commandName != "" {
		fmt.Printf("✅ Test file for command '%s' generated successfully\n", commandName)
	} else {
		fmt.Println("✅ General test file generated successfully")
	}

	return nil
}
