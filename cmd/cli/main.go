package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/cli/cmd/cli/internal/adder"
	"github.com/go-zoox/cli/cmd/cli/internal/generator"
	"github.com/go-zoox/cli/cmd/cli/internal/prompts"
	"github.com/go-zoox/cli/interactive"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:  "cli",
		Usage: "CLI framework tools and utilities",
	})

	app.Register("init", &cli.Command{
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

	app.Register("add", &cli.Command{
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

	app.Run()
}

func runInit(ctx *cli.Context) error {
	// 显示欢迎信息
	fmt.Println("🚀 CLI Scaffold - Generate CLI applications quickly!")
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

	// 收集 Flags
	var flags []prompts.FlagInfo
	if !skipInteractive {
		fmt.Println("\n📋 Let's add some flags to your CLI:")
		flags, err = prompts.CollectFlags()
		if err != nil {
			return err
		}
	}

	// 如果是多命令模式，收集命令信息
	var commands []prompts.CommandInfo
	if info.Type == "multiple" && !skipInteractive {
		fmt.Println("\n📦 Let's add commands to your CLI:")
		commands, err = prompts.CollectCommands()
		if err != nil {
			return err
		}
	}

	// 确认信息
	fmt.Println("\n📝 Project Summary:")
	fmt.Printf("  Name: %s\n", info.Name)
	fmt.Printf("  Type: %s\n", info.Type)
	fmt.Printf("  Flags: %d\n", len(flags))
	if info.Type == "multiple" {
		fmt.Printf("  Commands: %d\n", len(commands))
	}

	if !skipInteractive {
		confirm, err := prompts.Confirm("Proceed with generation?", true)
		if err != nil {
			return err
		}
		if !confirm {
			fmt.Println("Cancelled.")
			return nil
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
