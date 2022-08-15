package main

import (
	"fmt"

	"github.com/go-zoox/cli/interactive"
)

func main() {
	name, err := interactive.Password("Please type your password ?", &interactive.PasswordOptions{
		Required: true,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Your password is：", name)
}
