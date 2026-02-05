package template

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Template represents a CLI project template
type Template struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Path        string            `json:"path"`
	Type        string            `json:"type"` // "local", "github", "gitlab"
	Variables   map[string]string `json:"variables,omitempty"`
}

// Manager manages templates
type Manager struct {
	templatesDir string
}

// NewManager creates a new template manager
func NewManager() (*Manager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	templatesDir := filepath.Join(homeDir, ".cli", "templates")
	if err := os.MkdirAll(templatesDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create templates directory: %w", err)
	}

	return &Manager{
		templatesDir: templatesDir,
	}, nil
}

// List returns all available templates
func (m *Manager) List() ([]Template, error) {
	var templates []Template

	// Add built-in templates
	builtinTemplates := m.getBuiltinTemplates()
	templates = append(templates, builtinTemplates...)

	// Load custom templates
	customTemplates, err := m.loadCustomTemplates()
	if err != nil {
		return nil, err
	}
	templates = append(templates, customTemplates...)

	return templates, nil
}

// Add adds a new template
func (m *Manager) Add(name, path, templateType string) error {
	// Validate template path
	if templateType == "local" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return fmt.Errorf("template path does not exist: %s", path)
		}
	}

	template := Template{
		Name:        name,
		Path:        path,
		Type:        templateType,
		Description: fmt.Sprintf("Template from %s", path),
	}

	return m.saveTemplate(template)
}

// Remove removes a template
func (m *Manager) Remove(name string) error {
	// Check if it's a built-in template
	builtinTemplates := m.getBuiltinTemplates()
	for _, t := range builtinTemplates {
		if t.Name == name {
			return fmt.Errorf("cannot remove built-in template: %s", name)
		}
	}

	templateFile := filepath.Join(m.templatesDir, name+".json")
	if _, err := os.Stat(templateFile); os.IsNotExist(err) {
		return fmt.Errorf("template not found: %s", name)
	}

	return os.Remove(templateFile)
}

// Get retrieves a template by name
func (m *Manager) Get(name string) (*Template, error) {
	// Check built-in templates first
	builtinTemplates := m.getBuiltinTemplates()
	for _, t := range builtinTemplates {
		if t.Name == name {
			return &t, nil
		}
	}

	// Load custom template
	templateFile := filepath.Join(m.templatesDir, name+".json")
	data, err := os.ReadFile(templateFile)
	if err != nil {
		return nil, fmt.Errorf("template not found: %s", name)
	}

	var template Template
	if err := json.Unmarshal(data, &template); err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	return &template, nil
}

// getBuiltinTemplates returns built-in templates
func (m *Manager) getBuiltinTemplates() []Template {
	return []Template{
		{
			Name:        "basic",
			Description: "Basic CLI template (default)",
			Path:        "builtin:basic",
			Type:        "builtin",
		},
		{
			Name:        "api-client",
			Description: "API client CLI template",
			Path:        "builtin:api-client",
			Type:        "builtin",
		},
		{
			Name:        "file-manager",
			Description: "File management CLI template",
			Path:        "builtin:file-manager",
			Type:        "builtin",
		},
		{
			Name:        "server",
			Description: "Server management CLI template",
			Path:        "builtin:server",
			Type:        "builtin",
		},
		{
			Name:        "database",
			Description: "Database CLI template",
			Path:        "builtin:database",
			Type:        "builtin",
		},
		{
			Name:        "devops",
			Description: "DevOps CLI template",
			Path:        "builtin:devops",
			Type:        "builtin",
		},
	}
}

// loadCustomTemplates loads custom templates from disk
func (m *Manager) loadCustomTemplates() ([]Template, error) {
	var templates []Template

	entries, err := os.ReadDir(m.templatesDir)
	if err != nil {
		if os.IsNotExist(err) {
			return templates, nil
		}
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile(filepath.Join(m.templatesDir, entry.Name()))
		if err != nil {
			continue
		}

		var template Template
		if err := json.Unmarshal(data, &template); err != nil {
			continue
		}

		templates = append(templates, template)
	}

	return templates, nil
}

// saveTemplate saves a template to disk
func (m *Manager) saveTemplate(template Template) error {
	templateFile := filepath.Join(m.templatesDir, template.Name+".json")
	data, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal template: %w", err)
	}

	return os.WriteFile(templateFile, data, 0644)
}
