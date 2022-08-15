package interactive

import (
	"github.com/AlecAivazis/survey/v2"
)

// PasswordOptions is the options for Password.
type PasswordOptions struct {
	Required bool
}

// Password asks for a text input.
func Password(question string, opts ...*PasswordOptions) (string, error) {
	required := false
	if len(opts) > 0 && opts[0] != nil {
		required = opts[0].Required
	}

	q := &survey.Password{Message: question}

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
