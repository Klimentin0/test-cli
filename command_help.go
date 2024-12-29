package main

import (
	"fmt"
)

func helpPrompt() error {
	fmt.Println()
	fmt.Println("Welcome to my test prompt")
	fmt.Println("how to use:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("^_^")
	fmt.Println()
	return nil
}
