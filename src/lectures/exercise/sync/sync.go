//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

func inputAnalize() []string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\x13')
	wordsSlice := strings.Fields(text)
	return wordsSlice
}

type TotalNumber struct {
	totalNumber int
	sync.Mutex
}

func countLetters(s string, wg *sync.WaitGroup, total *TotalNumber) {
	defer wg.Done()
	var letters int
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters += 1
		}
	}
	total.Lock()
	total.totalNumber += letters
	total.Unlock()
}

func main() {
	var wg sync.WaitGroup
	newTotal := TotalNumber{}
	newAnalize := inputAnalize()
	for i := 0; i < len(newAnalize); i++ {
		wg.Add(1)
		go countLetters(newAnalize[i], &wg, &newTotal)
	}
	wg.Wait()
	fmt.Println("\nThe total number of letters is", newTotal.totalNumber)
}
