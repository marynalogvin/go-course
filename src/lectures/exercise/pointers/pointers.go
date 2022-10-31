//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

type SecurityTag bool
type itemsWithTags struct {
	item string
	tag  SecurityTag
}

// * Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag) {
	*tag = true
}
func deactivate(tag *SecurityTag) {
	*tag = false
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkOut(list []itemsWithTags) {
	for i := 0; i < len(list); i++ {
		deactivate(&list[i].tag)
	}
}

func main() {

	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	itemsList := []itemsWithTags{
		{"Pencil", Active},
		{"Pen", Active},
		{"Notebook", Active},
		{"Book", Active},
	}
	fmt.Println("Initial item's tags:", itemsList)

	// - Deactivate any one security tag in the array/slice
	deactivate(&itemsList[1].tag)
	fmt.Println("The second item is deactivated:", itemsList)

	// - Call the checkout() function to deactivate all tags
	checkOut(itemsList)
	fmt.Println("All items are deactivated:", itemsList)
	//  - Print out the array/slice after each change

}
