package adder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-zoox/cli/cmd/cli/internal/prompts"
)

// Adder adds commands or flags to existing CLI projects
type Adder struct {
	projectDir string
}

// NewAdder creates a new adder
func NewAdder(projectDir string) *Adder {
	return &Adder{
		projectDir: projectDir,
	}
}

// AddCommand adds a new command to an existing multiple commands CLI
func (a *Adder) AddCommand(cmd *prompts.CommandInfo) error {
	mainPath := filepath.Join(a.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", a.projectDir)
	}

	// Read the file
	src, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}

	content := string(src)

	// Check if it's a multiple commands CLI
	if !strings.Contains(content, "NewMultipleProgram") {
		return fmt.Errorf("this appears to be a single command CLI. Use 'cli init --type multiple' to create a multiple commands CLI first")
	}

	// Generate the new command code
	newCommandCode := a.generateCommandCode(cmd)

	// Find app.Run() and insert before it
	runIdx := strings.Index(content, "\n\tapp.Run()")
	if runIdx == -1 {
		runIdx = strings.Index(content, "\n\t\tapp.Run()")
	}
	if runIdx == -1 {
		return fmt.Errorf("could not find app.Run() in main.go")
	}

	// Insert the new command
	content = content[:runIdx] + newCommandCode + content[runIdx:]

	// Write back
	if err := os.WriteFile(mainPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %w", err)
	}

	return nil
}

// AddFlag adds a new flag to an existing CLI
func (a *Adder) AddFlag(flag *prompts.FlagInfo, commandName string) error {
	mainPath := filepath.Join(a.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", a.projectDir)
	}

	// For now, we'll use a simpler approach: read, modify as text, write
	src, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}

	content := string(src)

	// Generate flag code
	flagCode := a.generateFlagCode(flag)

	// Find where to insert the flag
	if commandName == "" {
		// Single command or global flags
		insertPos := a.findFlagsInsertPosition(content)
		if insertPos == -1 {
			return fmt.Errorf("could not find Flags array in main.go")
		}
		content = content[:insertPos] + flagCode + "\n\t\t" + content[insertPos:]
	} else {
		// Command-specific flags
		insertPos := a.findCommandFlagsInsertPosition(content, commandName)
		if insertPos == -1 {
			return fmt.Errorf("could not find command '%s' or its Flags array", commandName)
		}
		content = content[:insertPos] + flagCode + "\n\t\t\t" + content[insertPos:]
	}

	// Write back
	if err := os.WriteFile(mainPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %w", err)
	}

	return nil
}


func (a *Adder) generateCommandCode(cmd *prompts.CommandInfo) string {
	var flagsCode strings.Builder
	for _, flag := range cmd.Flags {
		flagsCode.WriteString(a.generateFlagCode(&flag))
		flagsCode.WriteString("\n\t\t\t")
	}

	flagsStr := flagsCode.String()
	if flagsStr != "" {
		flagsStr = "\n\t\tFlags: []cli.Flag{\n\t\t\t" + flagsStr + "\n\t\t},"
	}

	return fmt.Sprintf(`	app.Register("%s", &cli.Command{
		Name:  "%s",
		Usage: "%s",%s
		Action: func(ctx *cli.Context) error {
			// TODO: Implement %s command logic here
			fmt.Println("Running %s command")
			return nil
		},
	})

`, cmd.Name, cmd.Name, cmd.Usage, flagsStr, cmd.Name, cmd.Name)
}

func (a *Adder) generateFlagCode(flag *prompts.FlagInfo) string {
	typeFlag := capitalize(flag.Type) + "Flag"
	if flag.Type == "stringSlice" {
		typeFlag = "StringSliceFlag"
	}

	var parts []string
	parts = append(parts, fmt.Sprintf(`Name:  "%s"`, flag.Name))
	parts = append(parts, fmt.Sprintf(`Usage: "%s"`, flag.Usage))

	if flag.Default != "" && flag.Type != "bool" {
		if flag.Type == "string" || flag.Type == "stringSlice" {
			parts = append(parts, fmt.Sprintf(`Value: "%s"`, flag.Default))
		} else {
			parts = append(parts, fmt.Sprintf(`Value: %s`, flag.Default))
		}
	}

	if len(flag.Aliases) > 0 {
		aliases := make([]string, len(flag.Aliases))
		for i, alias := range flag.Aliases {
			aliases[i] = fmt.Sprintf(`"%s"`, alias)
		}
		parts = append(parts, fmt.Sprintf(`Aliases: []string{%s}`, strings.Join(aliases, ", ")))
	}

	if len(flag.EnvVars) > 0 {
		envVars := make([]string, len(flag.EnvVars))
		for i, envVar := range flag.EnvVars {
			envVars[i] = fmt.Sprintf(`"%s"`, envVar)
		}
		parts = append(parts, fmt.Sprintf(`EnvVars: []string{%s}`, strings.Join(envVars, ", ")))
	}

	return fmt.Sprintf(`&cli.%s{
				%s,
			}`, typeFlag, strings.Join(parts, ",\n\t\t\t\t"))
}

func (a *Adder) findFlagsInsertPosition(content string) int {
	// Find "Flags: []cli.Flag{" pattern
	idx := strings.Index(content, "Flags: []cli.Flag{")
	if idx == -1 {
		return -1
	}

	// Find the closing brace of the last flag before app.Run()
	runIdx := strings.Index(content[idx:], "app.Run()")
	if runIdx == -1 {
		// Find the closing brace of Flags array
		braceCount := 0
		started := false
		for i := idx; i < len(content); i++ {
			if content[i] == '{' {
				braceCount++
				started = true
			} else if content[i] == '}' {
				braceCount--
				if started && braceCount == 0 {
					return i
				}
			}
		}
		return -1
	}

	// Find the last flag before app.Run()
	flagsEnd := idx + runIdx
	for i := flagsEnd - 1; i >= idx; i-- {
		if content[i] == '}' && i+1 < len(content) && content[i+1] == ',' {
			return i + 2
		}
	}

	return idx + strings.Index(content[idx:], "{") + 1
}

func (a *Adder) findCommandFlagsInsertPosition(content string, commandName string) int {
	// Find the command registration
	pattern := fmt.Sprintf(`app.Register("%s"`, commandName)
	idx := strings.Index(content, pattern)
	if idx == -1 {
		return -1
	}

	// Find Flags array within this command
	flagsIdx := strings.Index(content[idx:], "Flags: []cli.Flag{")
	if flagsIdx == -1 {
		return -1
	}

	flagsStart := idx + flagsIdx
	// Find the closing brace of the last flag
	braceCount := 0
	started := false
	for i := flagsStart; i < len(content); i++ {
		if content[i] == '{' {
			braceCount++
			started = true
		} else if content[i] == '}' {
			braceCount--
			if started && braceCount == 0 {
				return i
			}
		}
	}

	return flagsStart + strings.Index(content[flagsStart:], "{") + 1
}


func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
