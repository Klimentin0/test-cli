package main

import (
	"fmt"
	"os"
)

func quitPrompt() error {
	fmt.Println("Exiting prompt")
	os.Exit(0)
	return nil
}
