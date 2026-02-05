package update

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Updater updates commands or flags in an existing CLI project.
type Updater struct {
	projectDir string
}

// NewUpdater creates updater.
func NewUpdater(projectDir string) *Updater {
	return &Updater{projectDir: projectDir}
}

func (u *Updater) read() (string, string, error) {
	mainPath := filepath.Join(u.projectDir, "main.go")
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to read main.go: %w", err)
	}
	return mainPath, string(data), nil
}

func (u *Updater) write(mainPath, content string) error {
	return os.WriteFile(mainPath, []byte(content), 0644)
}

// UpdateCommand updates command usage text.
func (u *Updater) UpdateCommand(name, newUsage string) error {
	mainPath, content, err := u.read()
	if err != nil {
		return err
	}

	pattern := fmt.Sprintf(`app.Register("%s"`, name)
	idx := strings.Index(content, pattern)
	if idx == -1 {
		return fmt.Errorf("command '%s' not found", name)
	}

	block := content[idx:]
	usagePattern := `Usage: "`
	usageIdx := strings.Index(block, usagePattern)
	if usageIdx == -1 {
		return fmt.Errorf("usage not found for command '%s'", name)
	}
	usageStart := idx + usageIdx + len(usagePattern)
	usageEnd := usageStart + strings.Index(content[usageStart:], `"`)
	if usageEnd <= usageStart {
		return fmt.Errorf("failed to locate usage end for '%s'", name)
	}

	newContent := content[:usageStart] + newUsage + content[usageEnd:]
	return u.write(mainPath, newContent)
}

// UpdateFlag updates flag usage/default within a command or global flags.
func (u *Updater) UpdateFlag(flagName, commandName, newUsage, newDefault string) error {
	mainPath, content, err := u.read()
	if err != nil {
		return err
	}

	var block string
	var blockStart int

	if commandName != "" {
		cmdPattern := fmt.Sprintf(`app.Register("%s"`, commandName)
		cmdStart := strings.Index(content, cmdPattern)
		if cmdStart == -1 {
			return fmt.Errorf("command '%s' not found", commandName)
		}
		closeIdx := strings.Index(content[cmdStart:], "\n\t})")
		if closeIdx == -1 {
			return fmt.Errorf("failed to find end of command '%s'", commandName)
		}
		cmdEnd := cmdStart + closeIdx + len("\n\t})")
		block = content[cmdStart:cmdEnd]
		blockStart = cmdStart
	} else {
		flagsIdx := strings.Index(content, "Flags: []cli.Flag{")
		if flagsIdx == -1 {
			return fmt.Errorf("no global Flags block found")
		}
		block = content[flagsIdx:]
		blockStart = flagsIdx
	}

	namePattern := fmt.Sprintf(`Name:  "%s"`, flagName)
	nameIdx := strings.Index(block, namePattern)
	if nameIdx == -1 {
		return fmt.Errorf("flag '%s' not found", flagName)
	}

	// Update usage
	if newUsage != "" {
		usagePattern := `Usage: "`
		usageRel := strings.Index(block[nameIdx:], usagePattern)
		if usageRel == -1 {
			return fmt.Errorf("usage for flag '%s' not found", flagName)
		}
		usageStart := blockStart + nameIdx + usageRel + len(usagePattern)
		usageEnd := usageStart + strings.Index(content[usageStart:], `"`)
		content = content[:usageStart] + newUsage + content[usageEnd:]
	}

	if newDefault != "" {
		defaultPattern := "Value:"
		blockUpdated := content[blockStart:]
		// Recompute because content may be changed
		nameIdxGlobal := blockStart + strings.Index(blockUpdated, namePattern)
		if nameIdxGlobal == -1 {
			return fmt.Errorf("flag '%s' not found for default update", flagName)
		}
		defRel := strings.Index(blockUpdated[strings.Index(blockUpdated, namePattern):], defaultPattern)
		if defRel == -1 {
			return fmt.Errorf("default value for flag '%s' not found", flagName)
		}
		defStart := blockStart + strings.Index(blockUpdated, namePattern) + defRel + len(defaultPattern)
		// find end of line
		defLineEnd := defStart + strings.Index(content[defStart:], "\n")
		if defLineEnd < defStart {
			defLineEnd = defStart + len(content[defStart:])
		}
		// replace value (keep spacing)
		content = content[:defStart] + " " + newDefault + content[defLineEnd:]
	}

	return u.write(mainPath, content)
}
