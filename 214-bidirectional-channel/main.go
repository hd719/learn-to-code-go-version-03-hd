package main

import "fmt"

func main() {
	c := make(chan int, 2)    // Create a channel of type int
	cr := make(chan<- int, 2) // Create a channel of type int (send only channel) fmt.Printf(<-c) will not work
	cs := make(<-chan int, 2) // Create a channel of type int (receive only channel) fmt.Printf(c <- 42) will not work

	c <- 42 // Send a first value to the channel
	c <- 43 // Send a second value to the buffered channel

	// Bidirectional channel (send values and pull values out of the channel)
	fmt.Println(<-c) // And now we can print the value
	fmt.Println(<-c) // And now we can print the value

	fmt.Println("----------------")
	fmt.Printf("%T\n", c)

	fmt.Println("----------------")
	// General to specific converts, single direction to bidirectional
	cr = c
	cs = c

	// This will not work
	// c = cr
	// c = cs
}
