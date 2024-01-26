package main

import "fmt"

func main() {
	c := make(chan int) // Create a channel of type int

	// send
	go foo(c)

	// receive
	// go bar(c) we remove the go keyword because we want to block the main routine until the value is pulled out
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c) // Pulling the value out
}
