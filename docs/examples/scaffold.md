# CLI Tool Examples

This page contains examples of using the CLI tool to generate CLI applications.

## Example 1: Simple Greeting CLI

Generate a simple CLI that greets users:

```bash
cli init --name greet --type single
```

Interactive prompts:
- Project name: `greet`
- Usage: `A simple greeting CLI`
- Version: `0.1.0`
- CLI type: `single`
- Add flag: `name` (String, required)
- Add flag: `formal` (Bool, default: false)

Generated code will include:

```go
app.Command(func(ctx *cli.Context) error {
    name := ctx.String("name")
    formal := ctx.Bool("formal")
    
    if formal {
        fmt.Printf("Good day, %s!\n", name)
    } else {
        fmt.Printf("Hello, %s!\n", name)
    }
    return nil
})
```

Usage:
```bash
go run main.go --name "Alice"
go run main.go --name "Bob" --formal
```

## Example 2: Todo List CLI

Generate a multiple commands CLI for managing todos:

```bash
cli init --name todo --type multiple
```

Add commands:
1. **list** command
   - Flag: `--all` (Bool) - Show all todos including completed
   - Flag: `--limit` (Int) - Limit number of results

2. **add** command
   - Flag: `--title` (String, required) - Todo title
   - Flag: `--priority` (String) - Priority level

3. **complete** command
   - Flag: `--id` (Int, required) - Todo ID to complete

Generated structure:
```go
app.Register("list", &cli.Command{
    // ...
})

app.Register("add", &cli.Command{
    // ...
})

app.Register("complete", &cli.Command{
    // ...
})
```

Usage:
```bash
todo list
todo add --title "Buy groceries" --priority "high"
todo complete --id 1
```

## Example 3: Server CLI with Configuration

Generate a CLI for managing a server:

```bash
cli init --name server --type single
```

Add flags:
- `--port` (Int, default: 8080, alias: `-p`)
- `--host` (String, default: "localhost")
- `--config` (String, env: `SERVER_CONFIG`)
- `--verbose` (Bool, alias: `-v`)

Generated code:
```go
Flags: []cli.Flag{
    &cli.IntFlag{
        Name:    "port",
        Usage:   "Server port",
        Value:   8080,
        Aliases: []string{"p"},
    },
    &cli.StringFlag{
        Name:    "host",
        Usage:   "Server host",
        Value:   "localhost",
    },
    &cli.StringFlag{
        Name:    "config",
        Usage:   "Config file path",
        EnvVars: []string{"SERVER_CONFIG"},
    },
    &cli.BoolFlag{
        Name:    "verbose",
        Usage:   "Enable verbose logging",
        Aliases: []string{"v"},
    },
}
```

Usage:
```bash
server --port 3000 --host 0.0.0.0
server -p 3000 -v
SERVER_CONFIG=/path/to/config server
```

## Example 4: File Operations CLI

Generate a CLI for file operations:

```bash
cli init --name fileops --type multiple
```

Add commands:
1. **copy** command
   - `--source` (String, required)
   - `--destination` (String, required)
   - `--recursive` (Bool)

2. **move** command
   - `--source` (String, required)
   - `--destination` (String, required)

3. **delete** command
   - `--path` (String, required)
   - `--force` (Bool)

## Example 5: API Client CLI

Generate a CLI for interacting with an API:

```bash
cli init --name apiclient --type single
```

Add flags:
- `--endpoint` (String, required, env: `API_ENDPOINT`)
- `--token` (String, required, env: `API_TOKEN`)
- `--method` (String, default: "GET")
- `--data` (String) - JSON data
- `--headers` (StringSlice) - Additional headers

Usage:
```bash
apiclient --endpoint https://api.example.com/users --method GET
apiclient --endpoint https://api.example.com/users --method POST --data '{"name":"John"}'
API_ENDPOINT=https://api.example.com API_TOKEN=xxx apiclient --method GET
```

## Best Practices

1. **Use meaningful names**: Choose descriptive project and command names
2. **Add helpful descriptions**: Write clear usage descriptions for flags and commands
3. **Set defaults**: Provide sensible default values for optional flags
4. **Use aliases**: Add short aliases for commonly used flags
5. **Environment variables**: Enable environment variable support for sensitive data
6. **Organize commands**: Group related commands logically in multiple commands mode
