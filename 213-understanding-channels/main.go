package main

import "fmt"

func main() {
	c := make(chan int) // Create a channel of type int

	// c <- 42 // This sends a value to the channel, however you cannot print this `fmt.Print(<-c)` because channels block
	// fmt.Print(<-c)

	// Buffered channels (stay away from them)
	// Successful
	d := make(chan int, 1) // Create a buffered channel of type int with a buffer size of 1 meaning it can hold 1 value without blocking

	// Unsuccessful attempt to send a value to the channel
	// d <- 42
	// d <- 45 // This sends a value to the channel, however you cannot print this `fmt.Print(<-c)` because of the buffer

	// To fix the blocking issue, we can use a goroutine
	go func() {
		c <- 42 // Send a value to the channel
		d <- 43 // Send a value to the buffered channel
	}()

	fmt.Println(<-c) // And now we can print the value
	fmt.Println(<-d) // And now we can print the value
}
