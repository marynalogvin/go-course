//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:

package main

import "fmt"

// - Products must include the price and the name
type Product struct {
	name  string
	price int
}

//   - Using an array, create a shopping list with enough room
//     for 4 products
//   - Print to the terminal:
//   - The last item on the list
//   - The total cost of the items
//   - The total number of items

func printInfo(shoppingList [4]Product) {
	var totalCost, totalItems int

	for i := 0; i < len(shoppingList); i++ {
		item := shoppingList[i]
		totalCost += item.price

		if item.name != "" {
			totalItems += 1
		}
	}
	fmt.Println("The last item on the list is", shoppingList[totalItems-1])
	fmt.Println("The total cost of the items", totalCost)
	fmt.Println("The total number of items", totalItems)
}

func main() {
	//* Insert 3 products into the array
	shoppingList := [4]Product{
		{"Milk", 5},
		{"Bread", 2},
		{"Eggs", 6},
	}
	printInfo(shoppingList)
	//   - Add a fourth product to the list and print out the
	//     information again
	shoppingList[3] = Product{"Oranges", 4}
	printInfo(shoppingList)

}
