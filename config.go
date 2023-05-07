package cli

import (
	"fmt"

	"github.com/go-zoox/config"
)

// LoadConfigOptions ...
type LoadConfigOptions struct {
	Required bool
	FlagKey  string
}

// LoadConfig loads the configuration by app name.
func LoadConfig(ctx *Context, cfg interface{}, opts ...*LoadConfigOptions) error {
	flagKeyX := "config"
	isRequired := false
	if len(opts) > 0 && opts[0] != nil {
		flagKeyX = opts[0].FlagKey
		isRequired = opts[0].Required
	}

	if ctx.String(flagKeyX) == "" {
		configName := "config.yml"
		if ctx.Command.Name != "" {
			configName = fmt.Sprintf("%s.yml", ctx.Command.Name)
		}

		// try to load from config, ignore error
		err := config.Load(cfg, &config.LoadOptions{
			AppName: ctx.App.Name,
			Name:    configName,
		})
		if err != nil {
			if isRequired {
				return err
			}

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
