package prompts

import (
	"github.com/go-zoox/cli/interactive"
)

// SelectOption represents a select option
type SelectOption = interactive.SelectOption

// Select prompts for a selection
func Select(prompt string, options []SelectOption, opts *interactive.SelectOptions) (string, error) {
	return interactive.Select(prompt, options, opts)
}

// Text prompts for text input
func Text(prompt string, opts *interactive.TextOptions) (string, error) {
	return interactive.Text(prompt, opts)
}
