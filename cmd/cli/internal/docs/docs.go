package docs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Generator generates documentation
type Generator struct {
	projectDir string
}

// NewGenerator creates a new docs generator
func NewGenerator(projectDir string) *Generator {
	return &Generator{
		projectDir: projectDir,
	}
}

// Generate generates documentation
func (g *Generator) Generate(outputDir, format string) error {
	mainPath := filepath.Join(g.projectDir, "main.go")
	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		return fmt.Errorf("main.go not found in %s", g.projectDir)
	}

	if outputDir == "" {
		outputDir = filepath.Join(g.projectDir, "docs")
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Read main.go to extract info
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return fmt.Errorf("failed to read main.go: %w", err)
	}

	content := string(data)

	// Simple extraction (in real implementation, use AST)
	projectName := extractProjectName(content)
	usage := extractUsage(content)

	// Generate markdown
	mdContent := g.generateMarkdown(projectName, usage, content)
	mdPath := filepath.Join(outputDir, "README.md")
	if err := os.WriteFile(mdPath, []byte(mdContent), 0644); err != nil {
		return fmt.Errorf("failed to write markdown: %w", err)
	}

	fmt.Printf("✅ Documentation generated in %s\n", outputDir)
	return nil
}

func (g *Generator) generateMarkdown(name, usage, code string) string {
	tmpl := "# {{.Name}}\n\n{{.Usage}}\n\n## Installation\n\n" +
		"```bash\n" +
		"go get {{.ModuleName}}\n" +
		"```\n\n" +
		"## Usage\n\n" +
		"```bash\n" +
		"{{.Name}} --help\n" +
		"```\n\n" +
		"## Commands\n\n{{.Commands}}\n\n## License\n\nMIT\n"

	data := map[string]interface{}{
		"Name":       name,
		"Usage":      usage,
		"ModuleName": name,
		"Commands":   "See `--help` for available commands",
	}

	var buf strings.Builder
	t := template.Must(template.New("docs").Parse(tmpl))
	t.Execute(&buf, data)
	return buf.String()
}

func extractProjectName(content string) string {
	// Simple extraction
	if strings.Contains(content, "NewSingleProgram") {
		if idx := strings.Index(content, "Name:"); idx != -1 {
			start := idx + 5
			end := strings.Index(content[start:], ",")
			if end != -1 {
				name := strings.Trim(content[start:start+end], `" `)
				return name
			}
		}
	}
	return "cli-app"
}

func extractUsage(content string) string {
	if idx := strings.Index(content, "Usage:"); idx != -1 {
		start := idx + 6
		end := strings.Index(content[start:], ",")
		if end != -1 {
			usage := strings.Trim(content[start:start+end], `" `)
			return usage
		}
	}
	return "A CLI application"
}

// ManGenerator generates man pages
type ManGenerator struct {
	projectDir string
}

// NewManGenerator creates a new man page generator
func NewManGenerator(projectDir string) *ManGenerator {
	return &ManGenerator{
		projectDir: projectDir,
	}
}

// Generate generates man page
func (m *ManGenerator) Generate() error {
	fmt.Println("✅ Man page generation is a placeholder")
	fmt.Println("Note: Full man page generation not yet implemented")
	return nil
}

// CompletionGenerator generates shell completion scripts
type CompletionGenerator struct {
	projectDir string
}

// NewCompletionGenerator creates a new completion generator
func NewCompletionGenerator(projectDir string) *CompletionGenerator {
	return &CompletionGenerator{
		projectDir: projectDir,
	}
}

// Generate generates completion script
func (c *CompletionGenerator) Generate(shell string) (string, error) {
	// Placeholder implementations
	switch shell {
	case "bash":
		return c.generateBash(), nil
	case "zsh":
		return c.generateZsh(), nil
	case "fish":
		return c.generateFish(), nil
	case "powershell":
		return c.generatePowerShell(), nil
	default:
		return "", fmt.Errorf("unsupported shell: %s", shell)
	}
}

func (c *CompletionGenerator) generateBash() string {
	return `# Bash completion script (placeholder)
# Install: source <(cli completion bash)
`
}

func (c *CompletionGenerator) generateZsh() string {
	return `# Zsh completion script (placeholder)
# Install: source <(cli completion zsh)
`
}

func (c *CompletionGenerator) generateFish() string {
	return `# Fish completion script (placeholder)
# Install: cli completion fish | source
`
}

func (c *CompletionGenerator) generatePowerShell() string {
	return `# PowerShell completion script (placeholder)
# Install: cli completion powershell | Out-String | Invoke-Expression
`
}
