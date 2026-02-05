package devtools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Builder builds CLI projects
type Builder struct {
	projectDir string
}

// NewBuilder creates a new builder
func NewBuilder(projectDir string) *Builder {
	return &Builder{
		projectDir: projectDir,
	}
}

// Build builds the project
func (b *Builder) Build(outputPath, platform string) error {
	mainPath := filepath.Join(b.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", b.projectDir)
	}

	var cmd *exec.Cmd
	if platform != "" {
		parts := splitPlatform(platform)
		if len(parts) != 2 {
			return fmt.Errorf("invalid platform format: %s (expected: os/arch)", platform)
		}
		os.Setenv("GOOS", parts[0])
		os.Setenv("GOARCH", parts[1])
		defer func() {
			os.Unsetenv("GOOS")
			os.Unsetenv("GOARCH")
		}()
	}

	if outputPath == "" {
		outputPath = filepath.Join(b.projectDir, filepath.Base(b.projectDir))
		if runtime.GOOS == "windows" {
			outputPath += ".exe"
		}
	}

	cmd = exec.Command("go", "build", "-o", outputPath, ".")
	cmd.Dir = b.projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build failed: %s", string(output))
	}

	fmt.Printf("✅ Build successful: %s\n", outputPath)
	return nil
}

func splitPlatform(platform string) []string {
	parts := make([]string, 0, 2)
	for _, part := range []string{"linux", "darwin", "windows", "freebsd"} {
		if len(platform) > len(part) && platform[:len(part)] == part {
			parts = append(parts, part)
			if len(platform) > len(part)+1 {
				parts = append(parts, platform[len(part)+1:])
			} else {
				parts = append(parts, "amd64")
			}
			return parts
		}
	}
	return []string{"linux", "amd64"}
}

// Runner runs CLI projects
type Runner struct {
	projectDir string
}

// NewRunner creates a new runner
func NewRunner(projectDir string) *Runner {
	return &Runner{
		projectDir: projectDir,
	}
}

// Run runs the project
func (r *Runner) Run(args []string) error {
	mainPath := filepath.Join(r.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", r.projectDir)
	}

	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = r.projectDir
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// Watcher watches files and rebuilds
type Watcher struct {
	projectDir string
}

// NewWatcher creates a new watcher
func NewWatcher(projectDir string) *Watcher {
	return &Watcher{
		projectDir: projectDir,
	}
}

// Watch watches for file changes and rebuilds
func (w *Watcher) Watch(build, test bool) error {
	// Simple implementation: just run build/test in a loop
	// In a real implementation, this would use file system events
	fmt.Println("Watching for file changes...")
	fmt.Println("Note: Full file watching not yet implemented. This is a placeholder.")

	if build {
		builder := NewBuilder(w.projectDir)
		if err := builder.Build("", ""); err != nil {
			return err
		}
	}

	if test {
		cmd := exec.Command("go", "test", "./...")
		cmd.Dir = w.projectDir
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Print(string(output))
			return err
		}
		fmt.Print(string(output))
	}

	return nil
}
