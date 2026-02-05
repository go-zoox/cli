package listing

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// CommandInfo holds parsed command data.
type CommandInfo struct {
	Name  string   `json:"name" yaml:"name"`
	Flags []string `json:"flags,omitempty" yaml:"flags,omitempty"`
}

// ProjectInfo holds aggregated project data.
type ProjectInfo struct {
	GlobalFlags []string      `json:"globalFlags,omitempty" yaml:"globalFlags,omitempty"`
	Commands    []CommandInfo `json:"commands,omitempty" yaml:"commands,omitempty"`
}

// List reads main.go and extracts commands/flags information.
func List(projectDir string) (*ProjectInfo, error) {
	mainPath := filepath.Join(projectDir, "main.go")
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read main.go: %w", err)
	}
	content := string(data)

	info := &ProjectInfo{}

	// Global flags (single command config)
	info.GlobalFlags = extractFlags(content, "Flags: []cli.Flag{", "\n\t})")

	// Commands
	cmdPattern := `app.Register("`
	search := content
	for {
		idx := strings.Index(search, cmdPattern)
		if idx == -1 {
			break
		}
		start := idx + len(cmdPattern)
		if start >= len(search) {
			break
		}
		rest := search[start:]
		endName := strings.Index(rest, `"`)
		if endName == -1 {
			break
		}
		name := rest[:endName]

		// Locate block end for this command
		blockStart := idx
		afterNameIdx := start + endName
		// Search for closing "\n\t})" after blockStart
		closeIdx := strings.Index(search[afterNameIdx:], "\n\t})")
		if closeIdx == -1 {
			closeIdx = len(search)
		} else {
			closeIdx = afterNameIdx + closeIdx + len("\n\t})")
		}
		block := search[blockStart:closeIdx]

		cmd := CommandInfo{
			Name:  name,
			Flags: extractFlags(block, "Flags: []cli.Flag{", "\n\t})"),
		}
		info.Commands = append(info.Commands, cmd)

		// Move cursor forward
		search = search[closeIdx:]
	}

	return info, nil
}

// FormatOutput formats project info.
func FormatOutput(info *ProjectInfo, format string) (string, error) {
	switch format {
	case "json":
		out, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			return "", err
		}
		return string(out), nil
	case "yaml", "yml":
		out, err := yaml.Marshal(info)
		if err != nil {
			return "", err
		}
		return string(out), nil
	default:
		// table-like text
		var b strings.Builder
		if len(info.GlobalFlags) > 0 {
			b.WriteString("Global Flags:\n")
			for _, f := range info.GlobalFlags {
				b.WriteString(fmt.Sprintf("  - %s\n", f))
			}
		}
		if len(info.Commands) > 0 {
			if len(info.GlobalFlags) > 0 {
				b.WriteString("\n")
			}
			b.WriteString("Commands:\n")
			for _, c := range info.Commands {
				b.WriteString(fmt.Sprintf("- %s\n", c.Name))
				if len(c.Flags) > 0 {
					for _, f := range c.Flags {
						b.WriteString(fmt.Sprintf("    * %s\n", f))
					}
				}
			}
		}
		if b.Len() == 0 {
			b.WriteString("No commands or flags found.\n")
		}
		return b.String(), nil
	}
}

func extractFlags(content, startMarker, endMarker string) []string {
	var flags []string
	start := strings.Index(content, startMarker)
	if start == -1 {
		return flags
	}
	end := strings.Index(content[start:], endMarker)
	if end == -1 {
		end = len(content)
	} else {
		end = start + end
	}
	block := content[start:end]

	// Find lines with Name:  "xxx"
	lines := strings.Split(block, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "Name:") {
			continue
		}
		// Name:  "xxx"
		namePart := strings.TrimPrefix(line, "Name:")
		namePart = strings.TrimSpace(namePart)
		if strings.HasPrefix(namePart, `"`) && strings.Contains(namePart[1:], `"`) {
			// Name includes quotes
			endIdx := strings.Index(namePart[1:], `"`)
			if endIdx != -1 {
				flags = append(flags, namePart[1:1+endIdx])
			}
		}
	}
	return flags
}
