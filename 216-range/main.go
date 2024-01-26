package main

import "fmt"

func main() {
	c := make(chan int) // Create a channel of type int

	// send
	go foo(c)

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c) // Close the channel to avoid a deadlock
}

// go run main.go and if you remove close(c)
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /Users/hameldesai/Developer/Go-Todd/216-range/main.go:12 +0xc0
// exit status 2
