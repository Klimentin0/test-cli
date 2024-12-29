package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Test prompt>>>")
		reader.Scan()

		strSlice := cleanInput(reader.Text())
		if len(strSlice) == 0 {
			continue
		}

		commandName := strSlice[0]

		fmt.Printf("Your command was: %s\n", commandName)
	}
}

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)
	for i, word := range strSlice {
		strSlice[i] = strings.ToLower(word)
	}
	return strSlice
}
