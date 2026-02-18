package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  calc [add|sub|mul|div] num1 num2")
		fmt.Println("  wordcount \"your text here\"")
		fmt.Println("  readfile filename.txt")
		return
	}

	command := os.Args[1]

	switch command {

	case "calc":
		handleCalculator()

	case "wordcount":
		handleWordCount()

	case "readfile":
		handleFileRead()

	default:
		fmt.Println("Unknown command:", command)
	}
}
