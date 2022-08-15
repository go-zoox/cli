package interactive

import (
	"fmt"

	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/text"
)

// TextOptions is the options for InquireText.
type TextOptions struct {
	Default  string
	Required bool
}

// Text asks for a text input.
func Text(question string, opts ...*TextOptions) (string, error) {
	defaultValue := ""
	required := false
	if len(opts) > 0 && opts[0] != nil {
		defaultValue = opts[0].Default
		required = opts[0].Required
	}

	if defaultValue != "" {
		question = fmt.Sprintf("%s [default: %s]", question, defaultValue)
	}

	input := inf.NewText(
		text.WithPrompt(question),
	)

	if err := input.Display(); err != nil {
		return "", err
	}

	value := input.Value()
	if value == "" {
		return defaultValue, nil
	}

	if required {
		if value == "" {
			return Text(question, opts...)
		}
	}

	return value, nil
}
