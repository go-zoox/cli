# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- **CLI Tool**: New CLI tool with `init` and `add` commands to quickly generate and extend CLI application templates
  - `cli init`: Initialize new CLI projects
    - Interactive mode for guided project setup
    - Non-interactive mode with flags for automation
    - Support for single command CLI generation
    - Support for multiple commands CLI generation
    - Automatic flag configuration with types (string, int, bool, stringSlice)
    - Support for flag aliases and environment variables
    - Automatic generation of project structure (main.go, go.mod, README.md, .gitignore)
    - Template-based code generation with customizable output
  - `cli add`: Add commands or flags to existing CLI projects
    - Add new commands to multiple commands CLI
    - Add new flags to existing commands or global flags
    - Interactive mode for guided addition
    - Non-interactive mode for automation
    - Automatic code insertion into existing main.go

- **Documentation**: Comprehensive documentation website built with VitePress
  - Getting started guide
  - Complete API reference
  - Interactive examples
  - Loading component examples
  - Scaffold tool documentation
  - GitHub Pages deployment via GitHub Actions

### Features

#### Scaffold Tool Features

- **Interactive Prompts**:
  - Project name, usage, and version collection
  - CLI type selection (single/multiple commands)
  - Flag configuration with type selection
  - Command configuration for multiple commands mode
  - Alias and environment variable support

- **Code Generation**:
  - Complete main.go with CLI setup
  - Proper flag definitions with all options
  - Command handlers with TODO comments
  - Example code for accessing flag values
  - Go module file generation
  - README.md with project information
  - Standard .gitignore for Go projects

- **Usage Modes**:
  - Interactive mode: Guided step-by-step setup
  - Non-interactive mode: Quick generation with flags
  - Customizable output directory

### Documentation

- Added scaffold tool guide at `/guide/scaffold`
- Added scaffold examples at `/examples/scaffold`
- Updated VitePress configuration to include scaffold documentation
- Added comprehensive usage examples and best practices

### Technical Details

#### CLI Tool Structure

```
cmd/cli/
├── main.go                    # CLI tool entry point
└── internal/
    ├── prompts/              # Interactive prompt modules
    │   ├── project.go        # Project information collection
    │   ├── flags.go          # Flag configuration collection
    │   └── commands.go      # Command configuration collection
    └── generator/            # Code generation modules
        └── generator.go      # Template-based code generator
```

#### Installation

```bash
# Build from source
go build -o cli ./cmd/cli

# Install globally
go install github.com/go-zoox/cli/cmd/cli@latest
```

#### Usage Examples

```bash
# Interactive mode
cli init

# Non-interactive mode
cli init --name myapp --type single --output ./myapp --skip-interactive
```

### Changed

- Enhanced documentation structure with new sections
- Improved developer experience with scaffold tool

### Fixed

- N/A (initial scaffold tool release)

---

## Previous Releases

[Previous changelog entries would be listed here]
