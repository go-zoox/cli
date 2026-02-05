package formatter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Formatter formats Go code
type Formatter struct {
	projectDir string
}

// NewFormatter creates a new formatter
func NewFormatter(projectDir string) *Formatter {
	return &Formatter{
		projectDir: projectDir,
	}
}

// Format formats the project using gofmt and goimports
func (f *Formatter) Format(checkOnly bool) error {
	mainPath := filepath.Join(f.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", f.projectDir)
	}

	// Run gofmt
	if err := f.runGofmt(mainPath, checkOnly); err != nil {
		return err
	}

	// Run goimports if available
	if f.hasGoimports() {
		if err := f.runGoimports(mainPath, checkOnly); err != nil {
			return err
		}
	}

	return nil
}

func (f *Formatter) runGofmt(filePath string, checkOnly bool) error {
	args := []string{"-l", filePath}
	if checkOnly {
		args = append(args, "-d")
	} else {
		args = []string{"-w", filePath}
	}

	cmd := exec.Command("gofmt", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if checkOnly {
			if len(output) > 0 {
				fmt.Print(string(output))
				return fmt.Errorf("code is not formatted")
			}
			return nil
		}
		return fmt.Errorf("gofmt failed: %s", string(output))
	}

	if checkOnly && len(output) > 0 {
		fmt.Print(string(output))
		return fmt.Errorf("code is not formatted")
	}

	return nil
}

func (f *Formatter) runGoimports(filePath string, checkOnly bool) error {
	args := []string{"-l", filePath}
	if checkOnly {
		args = append(args, "-d")
	} else {
		args = []string{"-w", filePath}
	}

	cmd := exec.Command("goimports", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		if checkOnly {
			if len(output) > 0 {
				fmt.Print(string(output))
				return fmt.Errorf("imports are not organized")
			}
			return nil
		}
		return fmt.Errorf("goimports failed: %s", string(output))
	}

	if checkOnly && len(output) > 0 {
		fmt.Print(string(output))
		return fmt.Errorf("imports are not organized")
	}

	return nil
}

func (f *Formatter) hasGoimports() bool {
	_, err := exec.LookPath("goimports")
	return err == nil
}
