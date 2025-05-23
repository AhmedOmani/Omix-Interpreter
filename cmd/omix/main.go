// Package main is the entry point for the Omix interpreter.
// It handles command-line arguments, file reading, and the main execution flow.
package main

import (
	"bufio"
	"fmt"
	"omix-interpreter/internal/lexer"
	"os"
)

type Omix struct {
	hadError bool
}

func NewOmix() *Omix {
	return &Omix{
		hadError: false,
	}
}

// runFile reads and executes the contents of a file.
// It returns an error if the file cannot be read or if there are any lexical errors.
func (o *Omix) runFile(path string) error {
	// Read the entire file into memory
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Convert bytes to string and run the code
	return o.run(string(bytes))
}

// runPrompt starts an interactive REPL (Read-Eval-Print Loop).
// It continuously reads input from the user and executes it.
func (o *Omix) runPrompt() {
	// Create a scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Print welcome message
	fmt.Println("Omix Interpreter")
	fmt.Println("Type 'exit' to quit")

	// Main REPL loop
	for {
		// Print prompt and read input
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		// Get the input line
		line := scanner.Text()

		// Check for exit command
		if line == "exit" {
			break
		}

		// Execute the input
		if err := o.run(line); err != nil {
			fmt.Println(err)
		}
	}
}

// run executes the given source code.
// It creates a new lexer, scans for tokens, and reports any errors.
func (o *Omix) run(source string) error {
	// Create a new lexer with the source code
	tokenizer := lexer.NewTokenizer(source)

	// Scan for tokens
	tokens := tokenizer.ScanTokens()

	// Check for lexical errors
	if tokenizer.HasErrors() {
		// Print all errors
		for _, err := range tokenizer.GetErrors() {
			fmt.Println(err)
		}
		return fmt.Errorf("lexical errors occurred")
	}

	// Print all tokens (for debugging)
	for _, token := range tokens {
		fmt.Println(token)
	}

	return nil
}

// main is the entry point of the program.
// It handles command-line arguments and determines whether to run a file or start the REPL.
func main() {
	// Get command-line arguments
	args := os.Args[1:]

	// Check number of arguments
	if len(args) > 1 {
		fmt.Println("Usage: omix [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		// Run the specified file
		omix := NewOmix()
		if err := omix.runFile(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(70)
		}
	} else {
		// Start the interactive REPL
		omix := NewOmix()
		omix.runPrompt()
	}
}

func (o *Omix) error(line int, message string) {
	o.report(line, "", message)
}

func (o *Omix) report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
	o.hadError = true
}
