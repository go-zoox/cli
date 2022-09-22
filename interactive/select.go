package interactive

import "github.com/AlecAivazis/survey/v2"

// SelectOptions is the options for a select.
type SelectOptions struct {
	Default string
}

// SelectOption is the option for Select.
type SelectOption struct {
	Label string
	Value string
}

// Select asks for a selection.
func Select(question string, options []SelectOption, opts ...*SelectOptions) (string, error) {
	var defaultValue string
	if len(opts) > 0 && opts[0] != nil {
		defaultValue = opts[0].Default
	}

	choices := []string{}
	optionsIndexLabel := map[string]string{}
	for _, opt := range options {
		choices = append(choices, opt.Label)
		optionsIndexLabel[opt.Label] = opt.Value

		if defaultValue != "" && defaultValue == opt.Value {
			defaultValue = opt.Label
		}
	}

	q := &survey.Select{
		Message: question,
		Options: choices,
	}
	if defaultValue != "" {
		q.Default = defaultValue
	}

	var label string
	err := survey.AskOne(q, &label)
	if err != nil {
		return "", err
	}

	return optionsIndexLabel[label], nil
}
