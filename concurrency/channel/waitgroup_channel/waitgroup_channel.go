package main

import (
	"fmt"
	"sync"
)

type User struct {
	firstName string
	lastName  string
}

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)

	// Create channel using make
	userChan := make(chan User)

	// goroutine with receive channel
	go func() {
		defer wg.Done()

		u := <-userChan
		fmt.Println("Hi", u.firstName, u.lastName)
		fmt.Println("Print done...")
	}()

	// goroutine with send channel
	go func() {
		defer wg.Done()

		// After line 35 then this goroutine will pause and go to run another goroutine
		userChan <- User{firstName: "SUpot", lastName: "Sukvaree"}

		fmt.Println("Assign done...")
	}()

	wg.Wait()
}
