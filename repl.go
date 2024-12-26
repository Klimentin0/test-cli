package main

import (
	"strings"
)

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)
	for i, word := range strSlice {
		strSlice[i] = strings.ToLower(word)
	}
	return strSlice
}
