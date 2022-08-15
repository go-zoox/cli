package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	name, err := interactive.Select(
		"What is your favorite city?",
		[]interactive.SelectOption{
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
		&interactive.SelectOptions{
			Default: "guangzhou",
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your favorite city: ", name)
}
