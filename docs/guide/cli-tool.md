# CLI Tool - Complete Guide

The CLI tool (`cli`) is a comprehensive toolkit for managing CLI projects built with the go-zoox/cli framework. It provides commands for project initialization, management, code quality, and more.

## Installation

```bash
go install github.com/go-zoox/cli/cmd/cli@latest
```

Or build from source:

```bash
go build -o cli ./cmd/cli
```

## Commands Overview

The CLI tool provides the following commands:

### Project Management
- [`init`](#init) - Initialize a new CLI project
- [`add`](#add) - Add commands or flags to existing project
- [`remove`](#remove) - Remove commands or flags from project
- [`list`](#list) - List commands and flags in project
- [`update`](#update) - Update command or flag properties
- [`validate`](#validate) - Validate project structure

### Templates
- [`template`](#template) - Manage project templates

### Code Quality
- [`format`](#format) - Format code with gofmt/goimports
- [`lint`](#lint) - Lint code with golangci-lint/go vet
- [`test`](#test) - Generate test files

### Project Enhancement
- [`upgrade`](#upgrade) - Upgrade CLI framework version
- [`migrate`](#migrate) - Migrate project structure
- [`config`](#config) - Manage CLI tool configuration

### Development Tools
- [`build`](#build) - Build CLI project
- [`run`](#run) - Run CLI project
- [`watch`](#watch) - Watch for file changes

### Documentation
- [`docs`](#docs) - Generate documentation
- [`man`](#man) - Generate man page
- [`completion`](#completion) - Generate shell completion scripts

### Utilities
- [`version`](#version) - Show version information

---

## init

Initialize a new CLI project.

### Usage

```bash
cli init [options]
```

### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--name` | `-n` | Project name | Required |
| `--type` | `-t` | CLI type (single\|multiple) | `single` |
| `--output` | `-o` | Output directory | `.` |
| `--skip-interactive` | `-s` | Skip interactive prompts | `false` |

### Examples

```bash
# Interactive mode
cli init

# Mixed mode
cli init --name myapp --type multiple

# Non-interactive mode
cli init --name myapp --type single --output ./myapp --skip-interactive
```

### Generated Examples

The `init` command automatically generates example code to help you get started:

**Single Command Mode** includes:
- `--name` flag (string, default: "World", alias: `-n`)
- `--verbose` flag (bool, alias: `-v`)

**Multiple Commands Mode** includes:
- `list` command with `--all` flag (bool, alias: `-a`)
- `create` command with `--name` flag (string, required, alias: `-n`)

You can modify these examples or use `cli add` to add more commands and flags.

### See Also

- [Getting Started with CLI Tool](/guide/scaffold)

---

## add

Add a command or flag to an existing CLI project.

### Usage

```bash
cli add [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--command` | `-c` | Command name to add |
| `--flag` | `-f` | Flag name to add |
| `--to-command` | | Add flag to specific command |
| `--dir` | `-d` | Project directory | `.` |

### Examples

```bash
# Add a command
cli add --command delete

# Add a flag
cli add --flag verbose

# Add flag to specific command
cli add --flag limit --to-command list
```

---

## remove

Remove a command or flag from an existing CLI project.

### Usage

```bash
cli remove [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--command` | `-c` | Command name to remove |
| `--flag` | `-f` | Flag name to remove |
| `--from-command` | | Remove flag from specific command |
| `--dir` | `-d` | Project directory | `.` |
| `--force` | | Skip confirmation | `false` |

### Examples

```bash
# Remove a command
cli remove --command delete

# Remove a flag
cli remove --flag verbose

# Remove flag from specific command
cli remove --flag limit --from-command list
```

---

## list

List commands and flags in an existing CLI project.

### Usage

```bash
cli list [options]
```

### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--dir` | `-d` | Project directory | `.` |
| `--format` | `-f` | Output format (table\|json\|yaml) | `table` |

### Examples

```bash
# List all commands and flags
cli list

# List in JSON format
cli list --format json

# List in YAML format
cli list --format yaml
```

---

## update

Update command or flag properties in an existing CLI project.

### Usage

```bash
cli update [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--command` | `-c` | Command name to update |
| `--flag` | `-f` | Flag name to update |
| `--usage` | | New usage/description for command |
| `--flag-usage` | | New usage for flag |
| `--flag-default` | | New default value for flag |
| `--from-command` | | Update flag in specific command |
| `--dir` | `-d` | Project directory | `.` |

### Examples

```bash
# Update command usage
cli update --command list --usage "List all items"

# Update flag default value
cli update --flag port --flag-default 3000

# Update flag in specific command
cli update --flag limit --from-command list --flag-default 10
```

---

## validate

Validate an existing CLI project structure and code.

### Usage

```bash
cli validate [options]
```

### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--dir` | `-d` | Project directory | `.` |
| `--fix` | | Attempt to fix issues (gofmt) | `false` |
| `--format` | `-f` | Output format (text\|json) | `text` |

### Examples

```bash
# Validate project
cli validate

# Validate and fix issues
cli validate --fix

# Validate with JSON output
cli validate --format json
```

---

## template

Manage CLI project templates.

### Usage

```bash
cli template <subcommand> [options]
```

### Subcommands

#### list

List all available templates.

```bash
cli template list
```

#### add

Add a new template.

```bash
cli template add --name <name> --path <path> [--type <type>]
```

**Options:**
- `--name`, `-n`: Template name (required)
- `--path`, `-p`: Template path (required)
- `--type`, `-t`: Template type (local\|github\|gitlab) (default: `local`)

**Examples:**

```bash
# Add local template
cli template add --name my-template --path ./templates/my-template

# Add GitHub template
cli template add --name github-template --path github:user/repo --type github
```

#### remove

Remove a template.

```bash
cli template remove --name <name>
```

**Options:**
- `--name`, `-n`: Template name to remove (required)

#### use

Use a template to initialize a project.

```bash
cli template use --name <name> [--output <dir>]
```

**Options:**
- `--name`, `-n`: Template name to use (required)
- `--output`, `-o`: Output directory (default: `.`)

**Examples:**

```bash
# Use built-in template
cli template use --name api-client --output ./my-api-client

# Use custom template
cli template use --name my-template --output ./my-project
```

### Built-in Templates

- `basic` - Basic CLI template (default)
- `api-client` - API client CLI template
- `file-manager` - File management CLI template
- `server` - Server management CLI template
- `database` - Database CLI template
- `devops` - DevOps CLI template

---

## format

Format Go code using gofmt and goimports.

### Usage

```bash
cli format [options]
```

### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--dir` | `-d` | Project directory | `.` |
| `--check` | `-c` | Check if code is formatted (don't modify) | `false` |

### Examples

```bash
# Format code
cli format

# Check formatting
cli format --check
```

---

## lint

Lint Go code using golangci-lint or go vet.

### Usage

```bash
cli lint [options]
```

### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--dir` | `-d` | Project directory | `.` |
| `--fix` | `-f` | Automatically fix issues | `false` |

### Examples

```bash
# Lint code
cli lint

# Lint and fix issues
cli lint --fix
```

---

## test

Generate test files for commands or flags.

### Usage

```bash
cli test [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--dir` | `-d` | Project directory | `.` |
| `--command` | `-c` | Generate test for specific command |

### Examples

```bash
# Generate general test file
cli test

# Generate test for specific command
cli test --command list
```

---

## upgrade

Upgrade CLI framework version.

### Usage

```bash
cli upgrade [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--dir` | `-d` | Project directory | `.` |
| `--to` | `-t` | Upgrade to specific version |
| `--check` | `-c` | Check for available upgrades |

### Examples

```bash
# Check for upgrades
cli upgrade --check

# Upgrade to latest version
cli upgrade

# Upgrade to specific version
cli upgrade --to v1.2.0
```

---

## migrate

Migrate project structure.

### Usage

```bash
cli migrate [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--dir` | `-d` | Project directory | `.` |
| `--to` | `-t` | Target structure (required) |

### Examples

```bash
# Migrate single command to multiple commands
cli migrate --to multiple
```

---

## config

Manage CLI tool configuration.

### Usage

```bash
cli config [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--key` | `-k` | Config key |
| `--value` | `-v` | Config value (for set) |

### Examples

```bash
# List all config
cli config

# Get config value
cli config --key output.format

# Set config value
cli config --key output.format --value json
```

---

## build

Build CLI project.

### Usage

```bash
cli build [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--dir` | `-d` | Project directory | `.` |
| `--output` | `-o` | Output binary path |
| `--platform` | `-p` | Target platform (e.g., linux/amd64) |

### Examples

```bash
# Build for current platform
cli build

# Build for specific platform
cli build --platform linux/amd64

# Build with custom output
cli build --output ./bin/myapp
```

---

## run

Run CLI project.

### Usage

```bash
cli run [options] [-- <args>]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--dir` | `-d` | Project directory | `.` |

### Examples

```bash
# Run project
cli run

# Run with arguments
cli run -- --help
cli run -- --name "World"
```

---

## watch

Watch for file changes and rebuild.

### Usage

```bash
cli watch [options]
```

### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--dir` | `-d` | Project directory | `.` |
| `--build` | `-b` | Rebuild on changes | `false` |
| `--test` | `-t` | Run tests on changes | `false` |

### Examples

```bash
# Watch and rebuild
cli watch --build

# Watch and run tests
cli watch --test

# Watch, rebuild and test
cli watch --build --test
```

---

## docs

Generate documentation.

### Usage

```bash
cli docs [options]
```

### Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--dir` | `-d` | Project directory | `.` |
| `--output` | `-o` | Output directory | `./docs` |
| `--format` | `-f` | Output format | `markdown` |

### Examples

```bash
# Generate documentation
cli docs

# Generate to custom directory
cli docs --output ./documentation
```

---

## man

Generate man page.

### Usage

```bash
cli man [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--dir` | `-d` | Project directory | `.` |

### Examples

```bash
# Generate man page
cli man
```

---

## completion

Generate shell completion script.

### Usage

```bash
cli completion --shell <shell> [options]
```

### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--shell` | `-s` | Shell type (bash\|zsh\|fish\|powershell) (required) |
| `--dir` | `-d` | Project directory | `.` |

### Examples

```bash
# Generate bash completion
cli completion --shell bash

# Generate zsh completion
cli completion --shell zsh

# Install bash completion
cli completion --shell bash > /etc/bash_completion.d/cli

# Install zsh completion
cli completion --shell zsh > ~/.zsh/completions/_cli
```

---

## version

Show version information.

### Usage

```bash
cli version
```

### Examples

```bash
cli version
```

Output:
```
cli tool version 0.1.0
CLI framework: github.com/go-zoox/cli
```

---

## Tips and Best Practices

1. **Use Mixed Mode**: Provide what you know via flags, let the tool ask for the rest
2. **Validate Before Committing**: Run `cli validate` before committing changes
3. **Format Code**: Use `cli format` to ensure consistent code style
4. **Use Templates**: Leverage built-in templates for common project types
5. **Version Control**: Always commit generated code to version control
6. **Test Generation**: Use `cli test` to scaffold test files
7. **Documentation**: Generate docs with `cli docs` for better project documentation

---

## Troubleshooting

### Command not found

Make sure you've installed the CLI tool:

```bash
go install github.com/go-zoox/cli/cmd/cli@latest
```

### Build errors

Run `go mod tidy` to ensure dependencies are resolved:

```bash
go mod tidy
```

### Template not found

Check available templates:

```bash
cli template list
```

### Validation errors

Use `--fix` flag to automatically fix some issues:

```bash
cli validate --fix
```

---

## See Also

- [Getting Started](/guide/getting-started)
- [CLI Framework Guide](/guide/)
- [Examples](/examples/)
