package main

import (
	"fmt"
	"runtime"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// Call function in the usual way
	// Running it synchronously
	f("From main")

	// New goroutine and execute it
	go f("From new goroutine")

	// We can start goroutine from anonymous function
	go func(from string) {
		fmt.Println(from)
	}("From anonymous func")

	fmt.Scanln()
	fmt.Println("done")
}
