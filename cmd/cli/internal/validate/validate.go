package validate

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

// Result holds validation findings.
type Result struct {
	Issues []string `json:"issues"`
}

// Validate checks basic CLI project structure. If fix is true, it will format main.go.
func Validate(projectDir string, fix bool) (*Result, error) {
	mainPath := filepath.Join(projectDir, "main.go")
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read main.go: %w", err)
	}
	content := string(data)

	var issues []string

	if !strings.Contains(content, "cli.NewSingleProgram") && !strings.Contains(content, "cli.NewMultipleProgram") {
		issues = append(issues, "CLI initialization not found (expect cli.NewSingleProgram or cli.NewMultipleProgram)")
	}

	if strings.Contains(content, "cli.NewMultipleProgram") && !strings.Contains(content, "app.Register(") {
		issues = append(issues, "MultipleProgram detected but no commands registered (app.Register)")
	}

	if !strings.Contains(content, "Flags: []cli.Flag{") {
		issues = append(issues, "No Flags block found; ensure flags are defined where needed")
	}

	if !strings.Contains(content, "app.Run()") {
		issues = append(issues, "app.Run() not found; CLI will not execute")
	}

	if fix {
		formatted, err := format.Source(data)
		if err != nil {
			issues = append(issues, fmt.Sprintf("gofmt failed: %v", err))
		} else {
			if err := os.WriteFile(mainPath, formatted, 0644); err != nil {
				return nil, fmt.Errorf("failed to write formatted main.go: %w", err)
			}
		}
	}

	return &Result{Issues: issues}, nil
}
