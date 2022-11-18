//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	sumNonBlankLines := 0
	sumCommands := 0

	i := true
	for i {
		input, inputErr := r.ReadString('\n')
		inputWord := strings.Trim(input, "\n")

		switch {
		//* When the user enters either "hello" or "bye", the program
		//  should respond with a message after pressing the enter key.
		case inputWord == "hello":
			fmt.Println("Hello from bufio")
			sumCommands += 1
			sumNonBlankLines += 1
		case inputWord == "bye":
			fmt.Println("Bye from bufio")
			sumCommands += 1
			sumNonBlankLines += 1
		case inputWord == "Q", inputWord == "q":
			//* Whenever the user types a "Q" or "q", the program should exit.
			fmt.Println("Exit")
			i = false
		default:
			inputWord = strings.TrimSpace(inputWord)
			if inputWord != "" {
				sumNonBlankLines += 1
			}
		}

		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}
	fmt.Printf("The number of non-blank lines entered: %v\n", sumNonBlankLines)
	fmt.Printf("The number of commands entered: %v\n", sumCommands)
}
