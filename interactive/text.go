package interactive

import (
	"github.com/AlecAivazis/survey/v2"
)

// TextOptions is the options for Text.
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

	q := &survey.Input{Message: question}
	if defaultValue != "" {
		q.Default = defaultValue
	}

	var value string
	err := survey.AskOne(q, &value, func(options *survey.AskOptions) error {
		if required {
			options.Validators = []survey.Validator{survey.Required}
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return value, nil
}
