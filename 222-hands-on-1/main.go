package main

import "fmt"

func main() {
	c := make(chan int, 1) // buffer channel

	// c <- 42 // Doesnt work

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
