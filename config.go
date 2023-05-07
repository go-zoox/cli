package cli

import (
	"github.com/go-zoox/config"
)

// LoadConfig loads the configuration by app name.
func LoadConfig(ctx *Context, cfg interface{}, flagKey ...string) error {
	flagKeyX := "config"
	if len(flagKey) > 0 && flagKey[0] != "" {
		flagKeyX = flagKey[0]
	}

	if ctx.String(flagKeyX) == "" {
		// try to load from config, ignore error
		err := config.Load(cfg, &config.LoadOptions{
			Name: ctx.App.Name,
		})
		if err != nil {
			if !config.IsNotFoundErr(err) {
				return err
			}
		}

		return nil
	}

	return config.Load(cfg, &config.LoadOptions{
		FilePath: ctx.String(flagKeyX),
	})
}
