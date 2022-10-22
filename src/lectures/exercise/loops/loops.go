//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:

//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	//* Print integers 1 to 50, except:
	for i := 0; i < 51; i++ {
		//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
			// - Print "Fizz" if the integer is divisible by 3
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			// - Print "Buzz" if the integer is divisible by 5
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}

}
