// Encoding videos with fan-out pattern (parallelism)

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10 // const is a keyword that creates a constant value
	wg.Add(goroutines)

	// For all the values that are in c1 channel, we are going to send them to c2 channel
	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				wg.Add(1)

				// This is the fan-out pattern
				// More importantly, we are creating a closure here
				// We are passing the value of v as an arg to the function below, and specified it as a param as v2
				// This creates different scopes for each value of v
				// v2 is scoped to the goroutine, while v is scoped to the for loop
				// If we didnt pass v as an arg to the function we would have variable shadowing
				go func(v2 int) {
					c2 <- timeConsumingWork(v2)
					wg.Done()
				}(v)
			}

			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
