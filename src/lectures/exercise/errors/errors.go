//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

// package timeparse
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}
type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

// * Example time string: 14:07:33
func split(input string) (*Time, error) {
	result := strings.Split(input, ":")
	if len(result) != 3 {
		return nil, &TimeParseError{"Invalid number of time components", input}
	} else {
		hour, err := strconv.Atoi(result[0])
		if err != nil {
			return nil, &TimeParseError{"Error parsing hour", input}
		}
		minute, err := strconv.Atoi(result[1])
		if err != nil {
			return nil, &TimeParseError{"Error parsing minute", input}
		}
		second, err := strconv.Atoi(result[2])
		if err != nil {
			return nil, &TimeParseError{"Error parsing second", input}
		}
		if hour > 23 || hour < 0 {
			return nil, &TimeParseError{"Inccorect hour", input}
		}
		if minute > 59 || minute < 0 {
			return nil, &TimeParseError{"Inccorect minute", input}
		}
		if second > 59 || second < 0 {
			return nil, &TimeParseError{"Inccorect second", input}
		}
		return &Time{hour, minute, second}, nil
	}
}

func main() {
	time1, err1 := split("n:20:40")
	time2, err2 := split("12:20:61")
	time3, err3 := split("12:20")

	fmt.Println(time1, err1, "\n", time2, err2, "\n", time3, err3)

}
