# Prompt-Powered CLI Toolkit: Building a Beginner’s Toolkit for Go (Golang)
## 1. Title & Objective

### Title:
Prompt-Powered CLI Toolkit: Building a Beginner-Friendly Command-Line Application in Go

### Objective:
The objective of this project is to learn the Go programming language from scratch by building a beginner-friendly Command-Line Interface (CLI) toolkit. The development process was guided by AI-assisted prompting, iterative debugging, and structured refinement. The end goal was to produce a modular CLI application demonstrating fundamental Go concepts.

##  2. Quick Summary of the Technology

Go is a statically typed, compiled programming language developed by Google in 2009. It was designed for simplicity, performance, and efficient concurrency handling.

Go is widely used in backend systems, cloud infrastructure, and DevOps tools. A real-world example is Docker, which is written in Go.

Go emphasizes:

Simple syntax

Fast compilation

Strong typing

Built-in concurrency support

## 3. System Requirements

Operating System: Windows 10/11

Go (version 1.20+ recommended)

VS Code or any code editor

Terminal (PowerShell or Command Prompt)

Git (optional for version control)

## 4. Installation & Setup Instructions
### Step 1: Install Go

Download Go from:
https://go.dev/dl/

Verify installation:
go version
### Step 2: Create Project Folder
mkdir cli-toolkit
cd cli-toolkit
### Step 3: Initialize Go Module
go mod init cli-toolkit
### Step 4: Run the Program
go run .

## 5. Minimal Working Example
### Example: Simple CLI Calculator Command
Command:
go run . calc add 10 5
Expected Output:
Result: 15

### Code Snippet (main.go structure)
package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
	case "calc":
		handleCalculator()
	}
}

This demonstrates:
CLI argument handling
Swich-based command routing
Modular function calls 

## 6. AI Prompt Journal

### Prompt 1:
Act as a patient Go programming instructor teaching a complete beginner. I am a student with no experience in Go and I want to create a very basic CLI program. Please explain step by step how to install Go, create a simple main.go file, and run it in the terminal. Organize your answer into setup, simple code example, and how to run it, and keep the explanation clear and easy to understand.
### Response Summary:
The AI explains how to install Go, create a simple project folder, write a basic main.go file, and run the program from the terminal. It breaks the process into clear steps (setup, code example, and execution), includes a simple working example, and provides beginner-friendly explanations for each part.
### Evaluation:
Helpful for understanding argument parsing, but required refinement to handle errors properly.

### Prompt 2:
Act as a Go programming instructor. I am new to Go and struggling with compiler errors while building a CLI toolkit. Explain common beginner errors, what they mean, and how to fix them. Structure your response into understanding errors, examples with fixes, and debugging tips.  
### Response Summary:
The AI explains Go compilation and common beginner errors, showing incorrect code alongside corrected versions. It teaches how to read error messages, trace issues, and fix them step by step. It also provides best practices for debugging CLI projects, all in a beginner-friendly, actionable style.
### Evaluation:
This significantly reduced debugging time and clarified Go’s compilation behavior.

### Prompt 3:
"How do I structure a Go project into multiple files?"
### Response Summary:
AI explained that all files must share the same package and that Go compiles all .go files in the folder.
### Evaluation:
Improved understanding of modular structure and separation of concerns.

#### Reflection on AI Usage

AI-assisted development improved productivity by:

Accelerating syntax learning

Reducing debugging time

Providing structured explanations

Helping refine modular design

However, prompts had to be refined multiple times to achieve precise and correct outputs.

## 7. Common Issues & Fixes
| Error                           | Cause                          | Solution                            |
| ------------------------------- | ------------------------------ | ----------------------------------- |
| expected 'package', found 'EOF' | File was empty                 | Added `package main`                |
| undefined function              | Function not yet created       | Created corresponding `.go` file    |
| invalid number conversion       | Non-integer input              | Used `strconv.Atoi` with validation |
| division by zero                | Invalid mathematical operation | Added zero-check condition          |

## 8. References
Official Go Documentation – https://go.dev/doc/
Go by Example – https://gobyexample.com/
