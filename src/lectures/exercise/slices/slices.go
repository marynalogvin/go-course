// --Summary:
//
//	Create a program to manage parts on an assembly line.
//
// --Requirements:
// * Using a slice, create an assembly line that contains type Part
// * Create a function to print out the contents of the assembly line
// * Perform the following:
//   - Print out the contents of the assembly line at each step
//
// --Notes:
// * Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printParts(assemblyLine []Part) {
	for i := 0; i < len(assemblyLine); i++ {
		element := assemblyLine[i]
		fmt.Println(element)
	}
}
func main() {
	//   - Create an assembly line having any three parts
	assemblyLine := []Part{"CleanDetails", "VarifyDetails", "AssembleDetails"}
	fmt.Println("3 parts:")
	printParts(assemblyLine)
	//   - Add two new parts to the line
	assemblyLine = append(assemblyLine, "PackDetails", "UnloadDetails")
	fmt.Println("\n Added 2 parts:")
	printParts(assemblyLine)
	//   - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("\n Sliced:")
	printParts(assemblyLine)

}
