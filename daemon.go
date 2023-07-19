package cli

import (
	"fmt"

	"github.com/go-zoox/fs"
	"github.com/go-zoox/logger"
	"github.com/sevlyar/go-daemon"
)

// DaemonOptions ...
type DaemonOptions struct {
	PidFile string
	LogFile string
	WorkDir string
}

// Daemon runs as a daemon
func Daemon(ctx *Context, fn func() error, opts ...*DaemonOptions) (err error) {
	appName := ctx.App.Name
	if ctx.Command.Name != "" {
		appName = fmt.Sprintf("%s_%s", ctx.App.Name, ctx.Command.Name)
	}

	pidFile := fmt.Sprintf("/tmp/gzcli.daemon.%s.pid", appName)
	logFile := fmt.Sprintf("/tmp/gzcli.daemon.%s.log", appName)
	workdir := fs.CurrentDir()
	if len(opts) > 0 && opts[0] != nil {
		if opts[0].PidFile != "" {
			pidFile = opts[0].PidFile
		}
		if opts[0].LogFile != "" {
			logFile = opts[0].LogFile
		}
		if opts[0].WorkDir != "" {
			workdir = opts[0].WorkDir
		}
	}

	cntxt := &daemon.Context{
		PidFileName: pidFile,
		PidFilePerm: 0644,
		LogFileName: logFile,
		LogFilePerm: 0640,
		WorkDir:     workdir,
		// Umask:       027,
		Args: []string{},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		return err
	}
	if d != nil {
		logger.Infof("daemon started(pid: %d, log: %s)", d.Pid, logFile)
		return
	}
	defer cntxt.Release()

	return fn()
}
