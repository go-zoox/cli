package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/go-zoox/cli/cmd/cli/internal/prompts"
)

// Generator generates CLI project files
type Generator struct {
	outputDir string
}

// NewGenerator creates a new generator
func NewGenerator(outputDir string) *Generator {
	return &Generator{outputDir: outputDir}
}

// GenerateSingleCommand generates a single command CLI project
func (g *Generator) GenerateSingleCommand(info *prompts.ProjectInfo, flags []prompts.FlagInfo) error {
	// 创建目录结构
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// 生成 main.go
	if err := g.generateMainFile(info, flags, "single"); err != nil {
		return err
	}

	// 生成 go.mod
	if err := g.generateGoMod(info); err != nil {
		return err
	}

	// 生成 README.md
	if err := g.generateREADME(info); err != nil {
		return err
	}

	// 生成 .gitignore
	if err := g.generateGitignore(); err != nil {
		return err
	}

	return nil
}

// GenerateMultipleCommands generates a multiple commands CLI project
func (g *Generator) GenerateMultipleCommands(info *prompts.ProjectInfo, commands []prompts.CommandInfo) error {
	// 创建目录结构
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// 生成 main.go
	if err := g.generateMainFileMultiple(info, commands); err != nil {
		return err
	}

	// 生成 go.mod
	if err := g.generateGoMod(info); err != nil {
		return err
	}

	// 生成 README.md
	if err := g.generateREADME(info); err != nil {
		return err
	}

	// 生成 .gitignore
	if err := g.generateGitignore(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateMainFile(info *prompts.ProjectInfo, flags []prompts.FlagInfo, cliType string) error {
	tmpl := `package main

import (
	"fmt"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "{{.Name}}",
		Usage:   "{{.Usage}}",
		Version: "{{.Version}}",
		Flags: []cli.Flag{
{{range .Flags}}			&cli.{{.TypeFlag}}{
				Name:  "{{.Name}}",
				Usage: "{{.Usage}}",
{{if .HasDefault}}				Value: {{.DefaultValue}},
{{end}}{{if .HasAliases}}				Aliases: []string{ {{range .Aliases}}"{{.}}"{{end}} },
{{end}}{{if .HasEnvVars}}				EnvVars: []string{ {{range .EnvVars}}"{{.}}"{{end}} },
{{end}}			},
{{end}}		},
	})

	app.Command(func(ctx *cli.Context) error {
		// TODO: Implement your command logic here
		fmt.Println("Hello from {{.Name}}!")
{{range .Flags}}{{if ne .Type "bool"}}		fmt.Printf("{{.Name}}: %v\n", ctx.{{.Getter}}("{{.Name}}"))
{{else}}		if ctx.Bool("{{.Name}}") {
			fmt.Println("{{.Name}} is enabled")
		}
{{end}}{{end}}		return nil
	})

	app.Run()
}
`

	data := g.buildTemplateData(info, flags, nil)
	return g.writeTemplate("main.go", tmpl, data)
}

func (g *Generator) generateMainFileMultiple(info *prompts.ProjectInfo, commands []prompts.CommandInfo) error {
	// main.go 只负责创建 app，并从 commands 目录中注册命令
	mainTmpl := `package main

import (
	"{{.ModuleName}}/commands"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:  "{{.Name}}",
		Usage: "{{.Usage}}",
	})

{{range .Commands}}	app.Register(commands.{{.FuncName}}())
{{end}}

	app.Run()
}
`

	// 为每个命令生成 commands/<name>.go（每个文件一个命令，返回 *cli.Command）
	commandFileTmpl := `package commands

import (
	"fmt"

	"github.com/go-zoox/cli"
)

// {{.FuncName}} returns the "{{.Name}}" command.
func {{.FuncName}}() *cli.Command {
	return &cli.Command{
		Name:  "{{.Name}}",
		Usage: "{{.Usage}}",
		Flags: []cli.Flag{
{{range .Flags}}			&cli.{{.TypeFlag}}{
				Name:  "{{.Name}}",
				Usage: "{{.Usage}}",
{{if .HasDefault}}				Value: {{.DefaultValue}},
{{end}}{{if .HasAliases}}				Aliases: []string{ {{range .Aliases}}"{{.}}"{{end}} },
{{end}}{{if .HasEnvVars}}				EnvVars: []string{ {{range .EnvVars}}"{{.}}"{{end}} },
{{end}}			},
{{end}}		},
		Action: func(ctx *cli.Context) error {
			// TODO: Implement {{.Name}} command logic here
			fmt.Println("Running {{.Name}} command")
{{range .Flags}}{{if ne .Type "bool"}}			fmt.Printf("{{.Name}}: %v\n", ctx.{{.Getter}}("{{.Name}}"))
{{else}}			if ctx.Bool("{{.Name}}") {
				fmt.Println("{{.Name}} is enabled")
			}
{{end}}{{end}}			return nil
		},
	}
}
`

	// 构建模板数据，包含 ModuleName 和命令函数名
	data := g.buildTemplateDataMultiple(info, commands)

	// 为每个命令增加 FuncName 字段（用于函数名）
	cmds, _ := data["Commands"].([]map[string]interface{})
	for i, c := range cmds {
		name, _ := c["Name"].(string)
		cmds[i]["FuncName"] = capitalize(name)
	}
	data["Commands"] = cmds

	// 写 main.go
	if err := g.writeTemplate("main.go", mainTmpl, data); err != nil {
		return err
	}

	// 创建 commands 目录
	commandsDir := filepath.Join(g.outputDir, "commands")
	if err := os.MkdirAll(commandsDir, 0755); err != nil {
		return fmt.Errorf("failed to create commands directory: %w", err)
	}

	// 为每个命令写独立文件 commands/<name>.go
	for _, c := range cmds {
		name, _ := c["Name"].(string)
		funcName, _ := c["FuncName"].(string)
		cmdData := map[string]interface{}{
			"Name":     name,
			"Usage":    c["Usage"],
			"Flags":    c["Flags"],
			"FuncName": funcName,
		}
		filename := fmt.Sprintf("commands/%s.go", name)
		if err := g.writeTemplate(filename, commandFileTmpl, cmdData); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) generateGoMod(info *prompts.ProjectInfo) error {
	tmpl := `module {{.ModuleName}}

go 1.18

require github.com/go-zoox/cli latest
`

	data := map[string]interface{}{
		"ModuleName": info.Name,
	}

	return g.writeTemplate("go.mod", tmpl, data)
}

func (g *Generator) generateREADME(info *prompts.ProjectInfo) error {
	tmpl := "# {{.Name}}\n\n{{.Usage}}\n\n## Installation\n\n```bash\ngo get {{.ModuleName}}\n```\n\n## Usage\n\n```bash\n{{.Name}} --help\n```\n\n## Development\n\n```bash\ngo run main.go\n```\n\n## License\n\nMIT\n"

	data := map[string]interface{}{
		"Name":       info.Name,
		"Usage":      info.Usage,
		"ModuleName": info.Name,
	}

	return g.writeTemplate("README.md", tmpl, data)
}

func (g *Generator) generateGitignore() error {
	content := `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
`

	return g.writeFile(".gitignore", content)
}

func (g *Generator) buildTemplateData(info *prompts.ProjectInfo, flags []prompts.FlagInfo, commands []prompts.CommandInfo) map[string]interface{} {
	processedFlags := make([]map[string]interface{}, len(flags))
	for i, flag := range flags {
		processedFlags[i] = g.processFlag(flag)
	}

	return map[string]interface{}{
		"Name":    info.Name,
		"Usage":   info.Usage,
		"Version": info.Version,
		"Flags":   processedFlags,
	}
}

func (g *Generator) buildTemplateDataMultiple(info *prompts.ProjectInfo, commands []prompts.CommandInfo) map[string]interface{} {
	processedCommands := make([]map[string]interface{}, len(commands))
	for i, cmd := range commands {
		processedFlags := make([]map[string]interface{}, len(cmd.Flags))
		for j, flag := range cmd.Flags {
			processedFlags[j] = g.processFlag(flag)
		}
		processedCommands[i] = map[string]interface{}{
			"Name":  cmd.Name,
			"Usage": cmd.Usage,
			"Flags": processedFlags,
		}
	}

	return map[string]interface{}{
		"Name":       info.Name,
		"Usage":      info.Usage,
		"Version":    info.Version,
		"ModuleName": info.Name,
		"Commands":   processedCommands,
	}
}

func (g *Generator) processFlag(flag prompts.FlagInfo) map[string]interface{} {
	typeFlag := capitalize(flag.Type) + "Flag"
	if flag.Type == "stringSlice" {
		typeFlag = "StringSliceFlag"
	}

	getter := "String"
	switch flag.Type {
	case "int":
		getter = "Int"
	case "bool":
		getter = "Bool"
	case "stringSlice":
		getter = "StringSlice"
	}

	hasDefault := flag.Default != ""
	var defaultValue string
	if hasDefault {
		if flag.Type == "string" || flag.Type == "stringSlice" {
			defaultValue = fmt.Sprintf(`"%s"`, flag.Default)
		} else {
			defaultValue = flag.Default
		}
	}

	return map[string]interface{}{
		"Name":         flag.Name,
		"Type":         flag.Type,
		"TypeFlag":     typeFlag,
		"Usage":        flag.Usage,
		"Default":      flag.Default,
		"HasDefault":   hasDefault,
		"DefaultValue": defaultValue,
		"Aliases":      flag.Aliases,
		"HasAliases":   len(flag.Aliases) > 0,
		"EnvVars":      flag.EnvVars,
		"HasEnvVars":   len(flag.EnvVars) > 0,
		"Getter":       getter,
	}
}

func (g *Generator) writeTemplate(filename, tmplContent string, data interface{}) error {
	tmpl, err := template.New(filename).Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	file, err := os.Create(filepath.Join(g.outputDir, filename))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func (g *Generator) writeFile(filename, content string) error {
	file, err := os.Create(filepath.Join(g.outputDir, filename))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// capitalize capitalizes the first letter of a string
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
