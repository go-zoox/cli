package interactive

import "github.com/AlecAivazis/survey/v2"

// ConfirmOptions is the options for Confirm.
type ConfirmOptions struct {
	Default bool
}

// Confirm asks for a confirmation.
func Confirm(question string, opts ...*ConfirmOptions) (bool, error) {
	var defaultValue bool
	if len(opts) > 0 && opts[0] != nil {
		defaultValue = opts[0].Default
	}

	q := &survey.Confirm{
		Message: question,
		Default: defaultValue,
	}

	var value bool
	err := survey.AskOne(q, &value)
	if err != nil {
		return false, err
	}

	return value, nil
}
