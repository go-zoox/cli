package migrate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Migrator migrates CLI projects
type Migrator struct {
	projectDir string
}

// NewMigrator creates a new migrator
func NewMigrator(projectDir string) *Migrator {
	return &Migrator{
		projectDir: projectDir,
	}
}

// MigrateToMultiple migrates a single command CLI to multiple commands CLI
func (m *Migrator) MigrateToMultiple() error {
	mainPath := filepath.Join(m.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", m.projectDir)
	}

	// Read main.go
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}

	content := string(data)

	// Check if already multiple commands
	if strings.Contains(content, "NewMultipleProgram") {
		return fmt.Errorf("project is already a multiple commands CLI")
	}

	// Simple migration: replace NewSingleProgram with NewMultipleProgram
	// In a real implementation, this would be more sophisticated
	newContent := strings.Replace(content, "NewSingleProgram", "NewMultipleProgram", -1)
	newContent = strings.Replace(newContent, "SingleProgramConfig", "MultipleProgramConfig", -1)

	// Backup original
	backupPath := mainPath + ".backup"
	if err := os.WriteFile(backupPath, data, 0644); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	// Write new content
	if err := os.WriteFile(mainPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %w", err)
	}

	fmt.Println("✅ Project migrated to multiple commands CLI")
	fmt.Printf("Backup saved to: %s\n", backupPath)
	return nil
}
