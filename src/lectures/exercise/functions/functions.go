// --Summary:
//
//	Use functions to perform basic operations and print some
//	information to the terminal.
//
// --Requirements:
package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greeting(personsName string) {
	fmt.Println("Hello", personsName)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func anyMessage() string {
	return "Hello from anyMessage function"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(n1, n2, n3 int) int {
	return n1 + n2 + n3
}

// * Write a function that returns any number
func anyNumber() int {
	return 101
}

// * Write a function that returns any two numbers
func anyTwoNumbers() (int, int) {
	return 102, 103
}

func main() {
	//* Call every function at least once
	//  * Print the result
	greeting("Alex")
	fmt.Println(anyMessage())
	// * Add three numbers together using any combination of the existing functions.
	firstNumber, secondNumber := anyTwoNumbers()
	addThreeNumbers := add(1, 2, 3) + firstNumber + secondNumber
	fmt.Println("The adding result is", addThreeNumbers)

}
