package main

import (
	"fmt"
	"os"
)

func handleFileRead() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: readfile filename.txt")
		return
	}

	filename := os.Args[2]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(data))
}
