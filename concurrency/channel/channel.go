package main

import (
	"fmt"
)

/*
 * Channels are the pipes that connect concurrent goroutine.
 * You can send value into channel from one goroutine
 * and receive those values into another goroutine.
 */

func main() {
	// Create new channel with make
	// The channel are typed by the value they convey
	message1 := make(chan string)
	message2 := make(chan string)

	// Send a value to channel using channel <- syntax
	// Try to send "I'm awesome" to the message channel from new goroutine
	go func() { message1 <- "I'm awesome" }()
	go func() { message2 <- "Try other way" }()

	go func(from string) {
		newMsg := <-message2
		fmt.Println(from, ":", newMsg)
	}("Value from message2")

	// Use <-channel syntax receives a value from the channel
	// msg will receive a value from another goroutine to main goroutine
	msg := <-message1
	fmt.Println("Value from message1 :", msg)

	fmt.Scanln()
	fmt.Println("Done")
}
