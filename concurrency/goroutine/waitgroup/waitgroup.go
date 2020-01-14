package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// Set the maximum number of CPUs
	runtime.GOMAXPROCS(1)

	// Create alias
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start GO routine")

	// First child GO routine
	go func() {
		defer wg.Done()

		fmt.Println("\ngoroutine #1 run...")

		// Add this line to tell GO runtime to swap to the second goroutine
		//time.Sleep(1 * time.Microsecond)
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	// Second child GO routine
	go func() {
		defer wg.Done()

		fmt.Println("\ngoroutine #2 run...")

		// Add this line to tell GO runtime to swap to the second goroutine
		//time.Sleep(1 * time.Microsecond)
		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\nProgram teminated")
}
