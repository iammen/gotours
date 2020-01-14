package main

import (
	"fmt"
	"runtime"
)

// Print a list of numbers that pass with channel
func printNumbersFromChannel(text string, ns chan int) {
	for n := range ns {
		fmt.Println(text, n)
	}
}

func main() {
	// Limit logical processor
	runtime.GOMAXPROCS(1)

	s := "Test string"

	// Create channel
	numbers := make(chan int)

	// Create first goroutine, pass string and channel
	go printNumbersFromChannel(s, numbers)

	// Create second goroutine
	go func() {
		for i := 0; i < 20; i++ {
			numbers <- i
		}

		close(numbers)
	}()

	fmt.Scanln()

}
