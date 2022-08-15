package intertive

import (
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/confirm"
	"github.com/fzdwx/infinite/emoji"
)

// Confirm asks for a confirmation.
func Confirm(question string, defaultValue ...bool) (bool, error) {
	defaultValueX := false
	if len(defaultValue) > 0 {
		defaultValueX = defaultValue[0]
	}

	opts := []confirm.Option{
		confirm.WithPrompt(question),
		confirm.WithDisplayHelp(),
		confirm.WithSymbol(emoji.Question),
	}

	if defaultValueX {
		opts = append(opts, confirm.WithDefaultYes())
	}

	c := inf.NewConfirm(opts...)

	ok, err := c.Display()
	if err != nil {
		return false, err
	}

	return ok, nil
}
