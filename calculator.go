package main

import (
	"fmt"
	"os"
	"strconv"
)

func handleCalculator() {

	if len(os.Args) < 5 {
		fmt.Println("Usage: calc [add|sub|mul|div] num1 num2")
		return
	}

	operation := os.Args[2]
	num1Str := os.Args[3]
	num2Str := os.Args[4]

	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Please provide valid numbers.")
		return
	}

	switch operation {
	case "add":
		fmt.Println("Result:", num1+num2)
	case "sub":
		fmt.Println("Result:", num1-num2)
	case "mul":
		fmt.Println("Result:", num1*num2)
	case "div":
		if num2 == 0 {
			fmt.Println("Error: Division by zero.")
			return
		}
		fmt.Println("Result:", num1/num2)
	default:
		fmt.Println("Unknown operation:", operation)
	}
}
