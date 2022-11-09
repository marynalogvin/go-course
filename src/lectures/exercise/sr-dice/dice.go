//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dice := 2
	sides := 2
	rolls := 2

	for roll := 1; roll <= rolls; roll++ {
		sum := 0
		for die := 1; die <= dice; die++ {
			sum += random(sides)
			fmt.Printf("Roll # %v, Die #%v : %v\n", roll, die, random(sides))
		}
		switch {
		case sum%2 == 0:
			fmt.Println("Even")
			if sum == 2 && dice == 2 {
				fmt.Println("Snake eyes")
			}
		case sum%2 != 0:
			fmt.Println("Odd")
			if sum == 7 {
				fmt.Println("Lucky 7")
			}
		}
		fmt.Printf("The sum is %v\n", sum)
	}
}
