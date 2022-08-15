package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	name, err := interactive.Text("What is your name ?", &interactive.TextOptions{
		Default:  "Zero",
		Required: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Your name is：", name)
}
