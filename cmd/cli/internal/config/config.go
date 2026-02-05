package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Manager manages CLI tool configuration
type Manager struct {
	configFile string
	config     map[string]interface{}
}

// NewManager creates a new config manager
func NewManager() (*Manager, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	configFile := filepath.Join(homeDir, ".clirc")
	manager := &Manager{
		configFile: configFile,
		config:     make(map[string]interface{}),
	}

	// Load existing config
	if err := manager.load(); err != nil {
		// Config file doesn't exist, that's okay
	}

	return manager, nil
}

// Get gets a config value
func (m *Manager) Get(key string) (interface{}, error) {
	value, ok := m.config[key]
	if !ok {
		return nil, fmt.Errorf("config key not found: %s", key)
	}
	return value, nil
}

// Set sets a config value
func (m *Manager) Set(key string, value interface{}) error {
	m.config[key] = value
	return m.save()
}

// List lists all config values
func (m *Manager) List() map[string]interface{} {
	return m.config
}

func (m *Manager) load() error {
	data, err := os.ReadFile(m.configFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &m.config)
}

func (m *Manager) save() error {
	data, err := json.MarshalIndent(m.config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	return os.WriteFile(m.configFile, data, 0644)
}
