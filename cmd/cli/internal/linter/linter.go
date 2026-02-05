package linter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Linter lints Go code
type Linter struct {
	projectDir string
}

// NewLinter creates a new linter
func NewLinter(projectDir string) *Linter {
	return &Linter{
		projectDir: projectDir,
	}
}

// Lint runs linter on the project
func (l *Linter) Lint(fix bool) error {
	mainPath := filepath.Join(l.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", l.projectDir)
	}

	// Try golangci-lint first
	if l.hasGolangciLint() {
		return l.runGolangciLint(fix)
	}

	// Fallback to go vet
	return l.runGoVet()
}

func (l *Linter) runGolangciLint(fix bool) error {
	args := []string{"run"}
	if fix {
		args = append(args, "--fix")
	}

	cmd := exec.Command("golangci-lint", args...)
	cmd.Dir = l.projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(output))
		return fmt.Errorf("golangci-lint found issues")
	}

	if len(output) > 0 {
		fmt.Print(string(output))
	}

	return nil
}

func (l *Linter) runGoVet() error {
	cmd := exec.Command("go", "vet", "./...")
	cmd.Dir = l.projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Print(string(output))
		return fmt.Errorf("go vet found issues")
	}

	return nil
}

func (l *Linter) hasGolangciLint() bool {
	_, err := exec.LookPath("golangci-lint")
	return err == nil
}
