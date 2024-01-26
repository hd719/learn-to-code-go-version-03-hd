package main

import "fmt"

func main() {
	c := make(chan int) // Create a channel of type int

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	// We have 10 goroutines sending 10 values to the channel, so we need to receive 100 values from the channel, which is why we itereate 100 times
	for k := 0; k < 100; k++ {
		println(k, <-c)
	}

	fmt.Print("About to exit")
}
