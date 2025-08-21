package main

import (
	"bufio"
	"fmt"
	"os"

	"glox/lox"
	"glox/scanner"
)

func main() {
	lox.HadError = false

	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}

}

func runFile(filename string) error {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	run(string(bytes))

	if lox.HadError {
		os.Exit(65)
	}

	return nil
}

func runPrompt() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")
		scanner.Scan()
		line := scanner.Text()

		if line == "" {
			break
		}

		run(line)
		lox.HadError = false
	}

	return nil
}

func run(source string) error {
	// tokens := strings.FieldsSeq(source) // For now scan tokens by space delim.

	scan := scanner.NewScanner(source)
	tokens := scan.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}

	return nil
}
