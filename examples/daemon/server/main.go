package main

import (
	"fmt"
	"os/exec"

	"github.com/go-zoox/cli/daemon"
	"github.com/go-zoox/fetch"
	"github.com/go-zoox/logger"
)

func main() {
	// app, err := daemon.NewServer("127.0.0.1:8080")
	// app, err := daemon.NewServer("http://127.0.0.1:8080")
	app, err := daemon.NewServer("unix:///tmp/go-zoox-cli-dameon.sock")
	if err != nil {
		logger.Errorf("failed to create server: %v", err)
		return
	}

	app.Command(&daemon.Command{
		Name:  "ls",
		Usage: "List directory contents",
		Flags: []daemon.Flag{
			{
				Name:  "l",
				Type:  "bool",
				Usage: "List files in the long format",
			},
			{
				Name:  "a",
				Type:  "bool",
				Usage: "List hidden files in the short format",
			},
		},
	}, func(args ...string) (response string, err error) {
		c := exec.Command("ls", args...)

		r, err := c.Output()
		if err != nil {
			return "", err
		}

		return string(r), nil
	})

	app.Command(&daemon.Command{
		Name:  "md5",
		Usage: "Get MD5 of the specified text",
	}, func(args ...string) (response string, err error) {
		if len(args) == 0 {
			return "", fmt.Errorf("text is required")
		}

		text := args[0]
		r, err := fetch.Get(fmt.Sprintf("https://httpbin.zcorky.com/md5/%s", text))
		if err != nil {
			return "", err
		}

		if !r.Ok() {
			return "", fmt.Errorf("%s", r.String())
		}

		return r.Get(text).String(), nil
	})

	app.Command(&daemon.Command{
		Name:  "uuid",
		Usage: "Get an UUID",
	}, func(args ...string) (response string, err error) {
		r, err := fetch.Get("https://httpbin.zcorky.com/uuid")
		if err != nil {
			return "", err
		}

		if !r.Ok() {
			return "", fmt.Errorf("%s", r.String())
		}

		return r.Get("uuid").String(), nil
	})

	if err = app.Run(); err != nil {
		logger.Errorf("failed to run server: %v", err)
		return
	}
}
