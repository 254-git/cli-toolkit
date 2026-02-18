package main

import (
	"fmt"
	"os"
	"strings"
)

func handleWordCount() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: wordcount \"your text here\"")
		return
	}

	text := os.Args[2]
	words := strings.Fields(text)

	fmt.Println("Word count:", len(words))
}
