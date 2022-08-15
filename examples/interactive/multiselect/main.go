package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	name, err := interactive.MultiSelect(
		"What is your favorite city?",
		[]interactive.MultiSelectOption{
			{
				Label: "Beijing",
				Value: "beijing",
			},
			{
				Label: "Shanghai",
				Value: "shanghai",
			},
			{
				Label: "Guangzhou",
				Value: "guangzhou",
			},
		},
		&interactive.MultiSelectOptions{
			// Default: []string{"guangzhou"},
			Required: true,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your favorite city: ", name)
}
