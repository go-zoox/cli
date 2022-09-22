package interactive

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
)

// MultiSelectOptions is the options for a multi select.
type MultiSelectOptions struct {
	Default  []string
	Required bool
}

// MultiSelectOption is the option for MultiSelect.
type MultiSelectOption struct {
	Label string
	Value string
}

// MultiSelect asks for a selection.
func MultiSelect(question string, options []MultiSelectOption, opts ...*MultiSelectOptions) ([]string, error) {
	var defaultValue []string
	required := false
	if len(opts) > 0 && opts[0] != nil {
		defaultValue = opts[0].Default
		required = opts[0].Required
	}

	choices := []string{}
	optionsIndexLabel := map[string]string{}
	optionsIndexValue := map[string]string{}
	for _, opt := range options {
		choices = append(choices, opt.Label)
		optionsIndexLabel[opt.Label] = opt.Value
		optionsIndexValue[opt.Value] = opt.Label
	}

	var defaultValueX []string
	if len(defaultValue) > 0 {
		for _, val := range defaultValue {
			if label, ok := optionsIndexValue[val]; ok {
				defaultValueX = append(defaultValueX, label)
			}
		}
	}

	q := &survey.MultiSelect{
		Message: question,
		Options: choices,
		Default: defaultValueX,
	}

	var labels []string
	err := survey.AskOne(q, &labels, func(options *survey.AskOptions) error {
		if required {
			options.Validators = []survey.Validator{func(ans interface{}) error {
				if len(ans.([]core.OptionAnswer)) == 0 {
					return errors.New("value is required")
				}

				return nil
			}}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	var values []string
	for _, label := range labels {
		values = append(values, optionsIndexLabel[label])
	}

	return values, nil
}
