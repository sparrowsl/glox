package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(os.Args, args)

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

	return run(string(bytes))
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
	}

	return nil
}

func run(source string) error {
	tokens := strings.FieldsSeq(source)
	// For now scan tokens by space delim.

	for token := range tokens {
		fmt.Println(token)
	}

	return nil
}
