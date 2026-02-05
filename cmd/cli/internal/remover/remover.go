package remover

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Remover deletes commands or flags from an existing CLI project.
type Remover struct {
	projectDir string
}

// NewRemover creates a new remover instance.
func NewRemover(projectDir string) *Remover {
	return &Remover{projectDir: projectDir}
}

// backup creates a backup copy of main.go before modifications.
func (r *Remover) backup(mainPath string) error {
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return err
	}
	return os.WriteFile(mainPath+".bak", data, 0644)
}

// RemoveCommand removes a command registration from main.go (multiple commands CLI).
func (r *Remover) RemoveCommand(name string) error {
	mainPath := filepath.Join(r.projectDir, "main.go")
	contentBytes, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}
	content := string(contentBytes)

	pattern := fmt.Sprintf(`app.Register("%s"`, name)
	start := strings.Index(content, pattern)
	if start == -1 {
		return fmt.Errorf("command '%s' not found", name)
	}

	// Find the end of the command block: the first occurrence of "\n\t})" after start.
	endSearch := content[start:]
	endRel := strings.Index(endSearch, "\n\t})")
	if endRel == -1 {
		return fmt.Errorf("failed to locate end of command '%s'", name)
	}
	end := start + endRel + len("\n\t})")

	// Include following newline if present.
	if end < len(content) && content[end] == '\n' {
		end++
	}

	if err := r.backup(mainPath); err != nil {
		return fmt.Errorf("failed to backup main.go: %w", err)
	}

	newContent := content[:start] + content[end:]
	if err := os.WriteFile(mainPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %w", err)
	}
	return nil
}

// RemoveFlag removes a flag from global flags (single command) or from a specific command.
// If commandName is empty, attempts to remove from the first Flags block (single command).
func (r *Remover) RemoveFlag(flagName, commandName string) error {
	mainPath := filepath.Join(r.projectDir, "main.go")
	contentBytes, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}
	content := string(contentBytes)

	// Limit search scope to a block.
	var block string
	var blockStart int

	if commandName != "" {
		cmdPattern := fmt.Sprintf(`app.Register("%s"`, commandName)
		cmdStart := strings.Index(content, cmdPattern)
		if cmdStart == -1 {
			return fmt.Errorf("command '%s' not found", commandName)
		}
		endRel := strings.Index(content[cmdStart:], "\n\t})")
		if endRel == -1 {
			return fmt.Errorf("failed to locate end of command '%s'", commandName)
		}
		cmdEnd := cmdStart + endRel + len("\n\t})")
		block = content[cmdStart:cmdEnd]
		blockStart = cmdStart
	} else {
		// First Flags block
		flagsIdx := strings.Index(content, "Flags: []cli.Flag{")
		if flagsIdx == -1 {
			return fmt.Errorf("no Flags block found")
		}
		block = content[flagsIdx:]
		blockStart = flagsIdx
	}

	namePattern := fmt.Sprintf(`Name:  "%s"`, flagName)
	nameIdx := strings.Index(block, namePattern)
	if nameIdx == -1 {
		return fmt.Errorf("flag '%s' not found", flagName)
	}

	// Find start of flag struct (preceding &cli.)
	structStartRel := strings.LastIndex(block[:nameIdx], "&cli.")
	if structStartRel == -1 {
		return fmt.Errorf("failed to locate start of flag '%s'", flagName)
	}

	// Find end of flag struct: look for "}," after name
	closeRel := strings.Index(block[nameIdx:], "},")
	if closeRel == -1 {
		return fmt.Errorf("failed to locate end of flag '%s'", flagName)
	}
	structEndRel := nameIdx + closeRel + len("},")

	// Expand to include trailing newline if present
	if structEndRel < len(block) && block[structEndRel] == '\n' {
		structEndRel++
	}

	absStart := blockStart + structStartRel
	absEnd := blockStart + structEndRel

	if err := r.backup(mainPath); err != nil {
		return fmt.Errorf("failed to backup main.go: %w", err)
	}

	newContent := content[:absStart] + content[absEnd:]
	if err := os.WriteFile(mainPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %w", err)
	}
	return nil
}
