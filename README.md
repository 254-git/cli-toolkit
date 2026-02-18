# Prompt-Powered CLI Toolkit in Go

## Project Description

This project is a beginner-friendly Command-Line Interface (CLI) toolkit built with Go. 
The toolkit was developed as part of an AI-assisted learning project where Go was learned from scratch using structured prompts and iterative refinement.

The goal of the project is to demonstrate:
- Basic Go syntax and structure
- CLI argument handling
- Modular file organization
- Error handling
- AI-assisted development workflow

---

## Technology Used

- Go (Golang)
- Standard Library Packages:
  - fmt
  - os
  - strconv
  - strings

Go is a statically typed, compiled programming language designed for simplicity and performance.

---

## Features

### 1. Calculator
Performs basic arithmetic operations:
- Addition
- Subtraction
- Multiplication
- Division

Example:
go run . calc add 10 5

Output:
Result: 15

---

### 2. Word Counter

Counts the number of words in a given sentence.

Example:
go run . wordcount "Go is simple and powerful"

Output:
Word count: 5

---

### 3. File Reader

Reads and displays the contents of a text file.

Example:
create a file named test.txt. add som words the run:
go run . readfile test.txt

---

## System Requirements

- Go installed (version 1.20+ recommended)
- Terminal (PowerShell, Command Prompt, or Bash)
- VS Code or any text editor

---

## Setup Instructions

1. Clone or download the project folder.
2. Open the folder in VS Code.
3. Initialize Go module (if not already initialized):
          go mod init cli-toolkit   //replace cli-toolkit with name of folder
          
4. Run the program:

---

## Project Structure

cli-toolkit/
│
├── main.go # CLI routing logic
├── calculator.go # Calculator functions
├── text.go # Word count logic
├── file.go # File reading functionality
├── go.mod # Go module file
└── README.md # Project documentation
├── doc/ # Toolkit document and AI prompt journal



---

## Error Handling

The program includes:
- Validation for missing arguments
- Validation for invalid numbers
- Division by zero protection
- File reading error handling

---

## Learning Reflection

This project was built using AI-assisted prompting. 
Prompts were refined iteratively to:
- Understand Go syntax
- Debug CLI errors
- Improve modular structure
- Implement input validation

The AI-assisted approach significantly improved development speed and debugging efficiency.

---

## Author

Precious John Mumbe



