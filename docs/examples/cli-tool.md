# CLI Tool Examples

This page contains comprehensive examples of using the CLI tool to manage CLI projects.

## Example 1: Complete Project Lifecycle

### Step 1: Initialize Project

```bash
cli init --name todo --type multiple
```

Interactive prompts:
- Usage: `A todo list manager`
- Version: `0.1.0`
- Add commands: `list`, `add`, `complete`, `delete`

### Step 2: Add Commands

```bash
# Add list command with flags
cli add --command list
# Follow prompts to add --all and --limit flags

# Add add command
cli add --command add
# Follow prompts to add --title and --priority flags

# Add complete command
cli add --command complete
# Follow prompts to add --id flag
```

### Step 3: Validate Project

```bash
cli validate
```

### Step 4: Format Code

```bash
cli format
```

### Step 5: List Commands

```bash
cli list
```

Output:
```
Commands:
  list     - List todos
  add      - Add a new todo
  complete - Complete a todo
  delete   - Delete a todo

Flags:
  Global:
    (none)
  
  Command 'list':
    --all    (bool)   Show all todos including completed
    --limit  (int)    Limit number of results
  
  Command 'add':
    --title     (string)  Todo title
    --priority  (string)  Priority level
```

### Step 6: Update Command

```bash
cli update --command list --usage "List all todos with optional filters"
```

### Step 7: Build and Run

```bash
# Build
cli build

# Run
cli run -- list --all
```

---

## Example 2: Using Templates

### List Available Templates

```bash
cli template list
```

Output:
```
Available templates:

  basic         - Basic CLI template (default) (builtin)
  api-client    - API client CLI template (builtin)
  file-manager  - File management CLI template (builtin)
  server        - Server management CLI template (builtin)
  database      - Database CLI template (builtin)
  devops        - DevOps CLI template (builtin)
```

### Use a Template

```bash
cli template use --name api-client --output ./my-api-client
```

### Add Custom Template

```bash
# Create template directory
mkdir -p ~/templates/my-template
# Add template files...

# Register template
cli template add --name my-template --path ~/templates/my-template
```

---

## Example 3: Code Quality Workflow

### Format Code

```bash
# Format all code
cli format

# Check formatting without modifying
cli format --check
```

### Lint Code

```bash
# Lint code
cli lint

# Lint and auto-fix
cli lint --fix
```

### Generate Tests

```bash
# Generate general test
cli test

# Generate test for specific command
cli test --command list
```

### Validate Project

```bash
# Validate
cli validate

# Validate and fix
cli validate --fix

# Validate with JSON output
cli validate --format json
```

---

## Example 4: Project Management

### Add Flag to Command

```bash
cli add --flag verbose --to-command list
```

Interactive prompts:
- Flag type: `bool`
- Usage: `Enable verbose output`
- Alias: `v`

### Update Flag

```bash
cli update --flag verbose --from-command list --flag-default true
```

### Remove Flag

```bash
cli remove --flag verbose --from-command list
```

### Remove Command

```bash
cli remove --command delete
```

---

## Example 5: Cross-Platform Build

### Build for Current Platform

```bash
cli build
```

### Build for Linux

```bash
cli build --platform linux/amd64
```

### Build for Multiple Platforms

```bash
# Linux
cli build --platform linux/amd64 --output ./bin/todo-linux-amd64

# macOS
cli build --platform darwin/amd64 --output ./bin/todo-darwin-amd64

# Windows
cli build --platform windows/amd64 --output ./bin/todo-windows-amd64.exe
```

---

## Example 6: Documentation Generation

### Generate Documentation

```bash
cli docs --output ./docs
```

### Generate Man Page

```bash
cli man
```

### Generate Shell Completions

```bash
# Bash
cli completion --shell bash > /etc/bash_completion.d/todo

# Zsh
cli completion --shell zsh > ~/.zsh/completions/_todo

# Fish
cli completion --shell fish > ~/.config/fish/completions/todo.fish
```

---

## Example 7: Project Migration

### Migrate Single to Multiple Commands

```bash
# Current: Single command CLI
# Target: Multiple commands CLI

cli migrate --to multiple
```

This will:
- Convert `NewSingleProgram` to `NewMultipleProgram`
- Create backup of original `main.go`
- Update project structure

---

## Example 8: Framework Upgrade

### Check for Upgrades

```bash
cli upgrade --check
```

### Upgrade to Latest

```bash
cli upgrade
```

### Upgrade to Specific Version

```bash
cli upgrade --to v1.2.0
```

---

## Example 9: Configuration Management

### List Configuration

```bash
cli config
```

### Get Configuration Value

```bash
cli config --key output.format
```

### Set Configuration Value

```bash
cli config --key output.format --value json
cli config --key template.default --value api-client
```

---

## Example 10: Development Workflow

### Watch and Rebuild

```bash
# Watch for changes and rebuild
cli watch --build
```

### Watch and Test

```bash
# Watch for changes and run tests
cli watch --test
```

### Run with Arguments

```bash
# Run project
cli run

# Run with specific arguments
cli run -- list --all --limit 10
```

---

## Best Practices

1. **Always validate before committing**:
   ```bash
   cli validate --fix
   ```

2. **Format code regularly**:
   ```bash
   cli format
   ```

3. **Use templates for common patterns**:
   ```bash
   cli template use --name api-client
   ```

4. **Generate tests early**:
   ```bash
   cli test --command list
   ```

5. **Keep documentation updated**:
   ```bash
   cli docs
   ```

6. **Use version control**:
   ```bash
   git add .
   git commit -m "Add new command"
   ```

---

## See Also

- [CLI Tool Guide](/guide/cli-tool)
- [Getting Started](/guide/getting-started)
- [CLI Framework Examples](/examples/)
