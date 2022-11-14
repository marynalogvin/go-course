//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Status int
type Book string

const (
	checkIn Status = iota
	checkOut
)

type LendingTime struct {
	checkOutTime string
	checkInTime  string
}

type BookStatusTime struct {
	status      Status
	lendingTime LendingTime
}

type Library struct {
	books map[Book]BookStatusTime
}

func (s BookStatusTime) String() string {
	switch s.status {
	case checkIn:
		return fmt.Sprintf("checkIn %v", s.lendingTime)
	case checkOut:
		return fmt.Sprintf("checkOut %v", s.lendingTime)
	default:
		return "unknown status"
	}
}
func printInfo(library *Library) {
	fmt.Printf("\nAll books and their statuses:\n")
	for i := range library.books {
		fmt.Printf("--------------\n %v: %v \n", i, library.books[i])
	}
	fmt.Printf("\n")

}

func bookAudit(library *Library, book Book, status Status) {
	time := time.Now()
	if status == checkOut {
		bookCheckOutTime := library.books[book]
		bookCheckOutTime.status = checkOut
		bookCheckOutTime.lendingTime.checkOutTime = (time.Format("01-02-2006 15:04:05"))
		library.books[book] = bookCheckOutTime
	} else if status == checkIn {
		bookCheckInTime := library.books[book]
		bookCheckInTime.status = checkIn
		bookCheckInTime.lendingTime.checkInTime = (time.Format("01-02-2006 15:04:05"))
		library.books[book] = bookCheckInTime
	}
}

func main() {
	libraryNew := map[Book]BookStatusTime{
		"Cinderella":     {},
		"Anna Karenina":  {},
		"The black swan": {},
		"Sapiens":        {},
	}
	printInfo(&Library{libraryNew})
	bookAudit(&Library{libraryNew}, "Cinderella", checkOut)
	printInfo(&Library{libraryNew})
	//delay to test the function bookAudit
	time.Sleep(5 * time.Second)
	bookAudit(&Library{libraryNew}, "Cinderella", checkIn)
	printInfo(&Library{libraryNew})
}
