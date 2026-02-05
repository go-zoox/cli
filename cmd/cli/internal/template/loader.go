package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Loader loads and applies templates
type Loader struct {
	templatesDir string
}

// NewLoader creates a new template loader
func NewLoader() (*Loader, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	templatesDir := filepath.Join(homeDir, ".cli", "templates")
	return &Loader{
		templatesDir: templatesDir,
	}, nil
}

// LoadTemplate loads a template by name
func (l *Loader) LoadTemplate(name string) (*TemplateData, error) {
	manager, err := NewManager()
	if err != nil {
		return nil, err
	}

	template, err := manager.Get(name)
	if err != nil {
		return nil, err
	}

	// Handle built-in templates
	if template.Type == "builtin" {
		return l.loadBuiltinTemplate(template.Name)
	}

	// Handle local templates
	if template.Type == "local" {
		return l.loadLocalTemplate(template.Path)
	}

	// Handle GitHub/GitLab templates (placeholder for future implementation)
	if strings.HasPrefix(template.Path, "github:") || strings.HasPrefix(template.Path, "gitlab:") {
		return nil, fmt.Errorf("remote templates not yet implemented")
	}

	return nil, fmt.Errorf("unsupported template type: %s", template.Type)
}

// TemplateData represents template data structure
type TemplateData struct {
	Files map[string]string // filename -> content
	Vars  map[string]string // template variables
}

// loadBuiltinTemplate loads a built-in template
func (l *Loader) loadBuiltinTemplate(name string) (*TemplateData, error) {
	// For now, built-in templates use the default generator
	// In the future, these can have custom templates
	return &TemplateData{
		Files: make(map[string]string),
		Vars:  make(map[string]string),
	}, nil
}

// loadLocalTemplate loads a template from local path
func (l *Loader) loadLocalTemplate(path string) (*TemplateData, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("template path does not exist: %s", absPath)
	}

	data := &TemplateData{
		Files: make(map[string]string),
		Vars:  make(map[string]string),
	}

	// Read template files
	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Skip hidden files and directories
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(absPath, path)
		if err != nil {
			return err
		}

		data.Files[relPath] = string(content)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk template directory: %w", err)
	}

	return data, nil
}
