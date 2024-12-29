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

		command, ok := getCommands()[commandName]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)
	for i, word := range strSlice {
		strSlice[i] = strings.ToLower(word)
	}
	return strSlice
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the repl",
			callback:    quitPrompt,
		},
		"help": {
			name:        "help",
			description: "help message",
			callback:    helpPrompt,
		},
	}
}
