# Daemon Mode

Daemon mode allows your CLI application to run in the background as a daemon process.

## Basic Usage

```go
package main

import (
	"log"
	"time"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:  "mydaemon",
		Usage: "A daemon application",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "daemon",
				Usage: "Run as daemon",
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		if ctx.Bool("daemon") {
			// Run as daemon
			return cli.Daemon(ctx, daemonFunc, &cli.DaemonOptions{
				PidFile: "/tmp/mydaemon.pid",
				LogFile: "/tmp/mydaemon.log",
			})
		}

		// Run normally
		return daemonFunc()
	})

	app.Run()
}

func daemonFunc() error {
	for {
		log.Println("Daemon is running...")
		time.Sleep(5 * time.Second)
	}
}
```

## Configuration

### DaemonOptions

| Field | Type | Description |
|-------|------|-------------|
| `PidFile` | `string` | Path to the PID file |
| `LogFile` | `string` | Path to the log file |
| `WorkDir` | `string` | Working directory for the daemon |

## Usage

```bash
# Run as daemon
./mydaemon --daemon

# Check if daemon is running
cat /tmp/mydaemon.pid

# View daemon logs
tail -f /tmp/mydaemon.log
```

## Complete Example

```go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:  "server",
		Usage: "A server daemon",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "daemon",
				Usage: "Run as daemon",
			},
			&cli.StringFlag{
				Name:  "pid-file",
				Usage: "PID file path",
				Value: "/tmp/server.pid",
			},
			&cli.StringFlag{
				Name:  "log-file",
				Usage: "Log file path",
				Value: "/tmp/server.log",
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		if ctx.Bool("daemon") {
			return cli.Daemon(ctx, serverFunc, &cli.DaemonOptions{
				PidFile: ctx.String("pid-file"),
				LogFile: ctx.String("log-file"),
			})
		}

		// Run in foreground
		return serverFunc()
	})

	app.Run()
}

func serverFunc() error {
	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start server
	log.Println("Server starting...")

	// Main loop
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Server is running...")
		case <-sigChan:
			log.Println("Server shutting down...")
			return nil
		}
	}
}
```
