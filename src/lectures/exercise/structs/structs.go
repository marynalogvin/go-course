//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.

//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

// * Create a rectangle structure containing its coordinates
type Coordinate struct {
	x, y int
}
type Rectangle struct {
	topLeft, bottomRight Coordinate
}

// - The functions must use the rectangle structure as the function parameter//
func width(rect Rectangle) int {
	return (rect.bottomRight.x - rect.topLeft.x)
}
func length(rect Rectangle) int {
	return (rect.topLeft.y - rect.bottomRight.y)
}

// * Using functions, calculate the area and perimeter of a rectangle,

func area(rect Rectangle) int {
	return length(rect) * width(rect)
}
func perimeter(rect Rectangle) int {
	return length(rect)*2 + width(rect)*2
}

// - Print the results to the terminal
func printInfo(rect Rectangle) {
	fmt.Println("The area is", area(rect))
	fmt.Println("The perimeter is", perimeter(rect))
}
func main() {
	Coordinate1 := Coordinate{0, 7}
	Coordinate2 := Coordinate{10, 0}
	Rectangle1 := Rectangle{Coordinate1, Coordinate2}
	//fmt.Println("Coordinates of rectangle are",Rectangle1)
	printInfo(Rectangle1)
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	//  - Print the new results to the terminal
	Rectangle1.topLeft.y *= 2
	Rectangle1.bottomRight.x *= 2
	printInfo(Rectangle1)

}
