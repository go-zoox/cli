package main

import (
	"github.com/go-zoox/cli/daemon"
	"github.com/go-zoox/logger"
)

func main() {
	// app, err := daemon.NewClient("http://127.0.0.1:8080")
	// app, err := daemon.NewClient("unix:///tmp/go-zoox-cli-dameon.sock")
	app, err := daemon.NewClient("unix:///tmp/go-zoox-cli-dameon.sock")
	if err != nil {
		logger.Errorf("failed to create client: %v", err)
		return
	}

	if err = app.Run(); err != nil {
		logger.Errorf("failed to run client: %v", err)
		return
	}
}
