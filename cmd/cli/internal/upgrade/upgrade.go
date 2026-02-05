package upgrade

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Upgrader upgrades CLI projects
type Upgrader struct {
	projectDir string
}

// NewUpgrader creates a new upgrader
func NewUpgrader(projectDir string) *Upgrader {
	return &Upgrader{
		projectDir: projectDir,
	}
}

// Check checks for available upgrades
func (u *Upgrader) Check() error {
	mainPath := filepath.Join(u.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", u.projectDir)
	}

	// Check go.mod for current version
	cmd := exec.Command("go", "list", "-m", "-versions", "github.com/go-zoox/cli")
	cmd.Dir = u.projectDir
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to check versions: %w", err)
	}

	fmt.Println("Current CLI framework version info:")
	fmt.Print(string(output))
	return nil
}

// Upgrade upgrades the project
func (u *Upgrader) Upgrade(version string) error {
	mainPath := filepath.Join(u.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", u.projectDir)
	}

	// Update go.mod
	var cmd *exec.Cmd
	if version != "" {
		cmd = exec.Command("go", "get", fmt.Sprintf("github.com/go-zoox/cli@%s", version))
	} else {
		cmd = exec.Command("go", "get", "-u", "github.com/go-zoox/cli")
	}
	cmd.Dir = u.projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to upgrade: %s", string(output))
	}

	fmt.Println("✅ Project upgraded successfully")
	fmt.Print(string(output))
	return nil
}
