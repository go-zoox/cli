package intertive

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	inf "github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components"
	"github.com/fzdwx/infinite/components/selection/singleselect"
)

// SelectOption is the option for InquireSelect.
type SelectOption[T any] struct {
	Key   string
	Value T
}

// Select asks for a selection.
func Select[T any](question string, options []SelectOption[T]) (*T, error) {
	choices := []string{}
	for _, opt := range options {
		choices = append(choices, opt.Key)
	}

	selected, err := inf.NewSingleSelect(
		choices,
		singleselect.WithDisableFilter(),
		singleselect.WithKeyBinding(components.SelectionKeyMap{
			Up: key.NewBinding(
				key.WithKeys("up"),
				key.WithHelp("↑", "move up"),
			),
			Down: key.NewBinding(
				key.WithKeys("down"),
				key.WithHelp("↓", "move down"),
			),
			Choice: key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "choose it"),
			),
			// @TODO
			Confirm: key.NewBinding(
				key.WithKeys("ctrl+c", "enter"),
				key.WithHelp("ctrl+c", "quit"),
			),
		}),
	).Display(question)
	if err != nil {
		return nil, fmt.Errorf("Quit")
	}

	return &options[selected].Value, nil
}
