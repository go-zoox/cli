package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	like, err := interactive.Confirm(
		"Do you like the book ?",
		&interactive.ConfirmOptions{
			Default: false,
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Your answer: ", like)
}
