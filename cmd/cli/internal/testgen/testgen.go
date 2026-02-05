package testgen

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// TestGenerator generates test files
type TestGenerator struct {
	projectDir string
}

// NewTestGenerator creates a new test generator
func NewTestGenerator(projectDir string) *TestGenerator {
	return &TestGenerator{
		projectDir: projectDir,
	}
}

// Generate generates test files for commands or flags
func (tg *TestGenerator) Generate(commandName string) error {
	mainPath := filepath.Join(tg.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", tg.projectDir)
	}

	if commandName != "" {
		return tg.generateCommandTest(commandName)
	}

	// Generate general test file
	return tg.generateGeneralTest()
}

func (tg *TestGenerator) generateCommandTest(commandName string) error {
	testFile := filepath.Join(tg.projectDir, fmt.Sprintf("%s_test.go", commandName))

	tmpl := `package main

import (
	"testing"
)

func Test{{.CommandName}}(t *testing.T) {
	// TODO: Implement test for {{.CommandName}} command
	t.Skip("Test not implemented yet")
}
`

	data := map[string]interface{}{
		"CommandName": strings.Title(commandName),
	}

	return tg.writeTemplate(testFile, tmpl, data)
}

func (tg *TestGenerator) generateGeneralTest() error {
	testFile := filepath.Join(tg.projectDir, "main_test.go")

	tmpl := `package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// TODO: Implement main test
	t.Skip("Test not implemented yet")
}
`

	return tg.writeTemplate(testFile, tmpl, nil)
}

func (tg *TestGenerator) writeTemplate(filePath, tmplContent string, data interface{}) error {
	tmpl, err := template.New(filepath.Base(filePath)).Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
