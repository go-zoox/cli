package cli

import (
	"fmt"

	ucli "github.com/urfave/cli/v2"
)

func init() {
	// @TODO Override the default version printer.
	ucli.VersionPrinter = func(c *Context) {
		_, _ = fmt.Fprintf(c.App.Writer, "v%s\n", c.App.Version)
	}
}

// Context ...
type Context = ucli.Context

// Command ...
type Command = ucli.Command

// Action ...
type Action = func(*Context) error

// type Action = ucli.ActionFunc

// Flag ...
type Flag = ucli.Flag

// StringFlag ...
type StringFlag = ucli.StringFlag

// Int64Flag ...
type Int64Flag = ucli.Int64Flag

// Float64Flag ...
type Float64Flag = ucli.Float64Flag

// IntFlag ...
type IntFlag = ucli.IntFlag

// BoolFlag ...
type BoolFlag = ucli.BoolFlag

// StringSliceFlag ...
type StringSliceFlag = ucli.StringSliceFlag

// IntSliceFlag ...
type IntSliceFlag = ucli.IntSliceFlag

// Int64SliceFlag ...
type Int64SliceFlag = ucli.Int64SliceFlag

// Float64SliceFlag ...
type Float64SliceFlag = ucli.Float64SliceFlag

// PathFlag ...
type PathFlag = ucli.PathFlag

// UintFlag ...
type UintFlag = ucli.UintFlag

// Uint64Flag ...
type Uint64Flag = ucli.Uint64Flag

// UintSliceFlag ...
type UintSliceFlag = ucli.UintSliceFlag

// Uint64SliceFlag ...
type Uint64SliceFlag = ucli.Uint64SliceFlag
