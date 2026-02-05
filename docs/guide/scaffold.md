# CLI Tool - Init Command

The `cli init` command helps you quickly generate CLI application templates using the go-zoox/cli framework.

## Installation

The CLI tool is included in the CLI framework. To use it, you need to build it:

```bash
go build -o cli ./cmd/cli
```

Or install it globally:

```bash
go install github.com/go-zoox/cli/cmd/cli@latest
```

## Quick Start

### Interactive Mode

The easiest way to use the CLI tool is in interactive mode:

```bash
cli init
```

This will guide you through:
1. Project name and description
2. CLI type (single or multiple commands)
3. Flags configuration
4. Commands configuration (for multiple commands mode)

### Mixed Mode (Recommended)

You can provide some parameters via command line, and the tool will only ask for the missing ones:

```bash
# Provide name, tool will ask for other details
cli init --name myapp

# Provide name and type, tool will ask for flags and commands
cli init --name myapp --type multiple

# Provide all parameters, tool will only ask for confirmation
cli init --name myapp --type single --output ./myapp
```

### Non-Interactive Mode

Use `--skip-interactive` to skip all prompts (requires all parameters):

```bash
cli init --name myapp --type single --output ./myapp --skip-interactive
```

## Usage

### Initialize a Single Command CLI

```bash
cli init --name myapp --type single
```

This creates a new single command CLI application with:
- `main.go` - Main application file
- `go.mod` - Go module file
- `README.md` - Project documentation
- `.gitignore` - Git ignore file

### Initialize a Multiple Commands CLI

```bash
cli init --name myapp --type multiple
```

This creates a new multiple commands CLI application where you can add multiple subcommands.

## Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--name` | `-n` | Project name | Required |
| `--type` | `-t` | CLI type (single\|multiple) | `single` |
| `--output` | `-o` | Output directory | `.` (current directory) |
| `--skip-interactive` | `-s` | Skip interactive prompts | `false` |

## Interactive Prompts

### Project Information

- **Project name**: The name of your CLI application
- **Usage/Description**: A brief description of what your CLI does
- **Version**: Version number (default: `0.1.0`)
- **CLI type**: Choose between single command or multiple commands

### Flags Configuration

For each flag, you'll be asked:
- **Flag name**: The name of the flag (e.g., `port`)
- **Flag type**: Choose from:
  - String
  - Int
  - Bool
  - String Slice
- **Usage/Description**: Description of what the flag does
- **Default value**: Optional default value
- **Short alias**: Optional short form (e.g., `-p` for `--port`)
- **Environment variable**: Optional environment variable name

### Commands Configuration (Multiple Commands Mode)

For each command, you'll be asked:
- **Command name**: The name of the command (e.g., `list`, `create`)
- **Usage/Description**: Description of what the command does
- **Flags**: Optional flags for this command (same process as above)

## Generated Files

### main.go

The generated `main.go` file includes:
- Complete CLI setup with your configuration
- Flag definitions
- Command handlers with TODO comments
- Example code showing how to access flag values

### go.mod

The generated `go.mod` file includes:
- Module name
- Go version (1.18)
- Dependency on `github.com/go-zoox/cli`

### README.md

A basic README with:
- Project name and description
- Installation instructions
- Usage examples

### .gitignore

A standard Go `.gitignore` file that excludes:
- Binaries
- Test files
- IDE files
- OS-specific files

## Examples

### Example 1: Simple Single Command CLI

```bash
scaffold init --name greet --type single
```

Then follow the prompts to add a `name` flag. The generated code will look like:

```go
app.Command(func(ctx *cli.Context) error {
    name := ctx.String("name")
    fmt.Printf("Hello, %s!\n", name)
    return nil
})
```

### Example 2: Multiple Commands CLI

```bash
scaffold init --name todo --type multiple
```

Add commands like `list`, `add`, `remove` with their respective flags.

## Add Command

The `cli add` command allows you to add new commands or flags to an existing CLI project.

### Add a Command

To add a new command to a multiple commands CLI:

```bash
# Fully interactive mode
cli add

# Provide command name, tool will ask for usage and flags
cli add --command mycommand
```

### Add a Flag

To add a new flag to your CLI:

```bash
# Fully interactive mode
cli add

# Provide flag name, tool will ask for type and other details
cli add --flag port

# Add flag to a specific command
cli add --flag port --to-command list
```

### Mixed Mode for Add Command

The `add` command supports mixed mode - you can provide partial information and the tool will ask for the rest:

```bash
# Provide command name, tool asks for usage and flags
cli add --command delete

# Provide flag name, tool asks for type, usage, etc.
cli add --flag verbose

# Provide flag name and target command, tool asks for flag details
cli add --flag limit --to-command list
```

## Next Steps

After generating your CLI project:

1. **Navigate to the project directory**:
   ```bash
   cd myapp
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run your CLI**:
   ```bash
   go run main.go
   ```

4. **Build your CLI**:
   ```bash
   go build -o myapp
   ```

5. **Test with flags**:
   ```bash
   ./myapp --help
   ./myapp --name "World"
   ```

## Tips

- Use mixed mode for the best experience - provide what you know, let the tool ask for the rest
- You can combine command-line flags with interactive prompts
- Use `--skip-interactive` only when you have all required parameters
- You can always edit the generated code to customize it
- The generated code includes TODO comments to guide you
- All flags support environment variables automatically
- Short aliases make your CLI more user-friendly

## Troubleshooting

### "command not found: cli"

Make sure you've built or installed the CLI tool:
```bash
go install github.com/go-zoox/cli/cmd/cli@latest
```

### "failed to create directory"

Make sure you have write permissions in the output directory.

### Generated code doesn't compile

Run `go mod tidy` to ensure all dependencies are properly resolved.
