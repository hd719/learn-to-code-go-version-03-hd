// https://go.dev/blog/pipelines

package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

	// Fan Out
	// Distribute the sq work across 2 goroutines that both read from in
	c1 := sq(in)
	c2 := sq(in)

	// Fan In
	// Consume the merged output from multiple channels
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	fmt.Printf("Type of cs: %T\n", cs)

	// We increment the WaitGroup before starting each goroutine to process an input channel. This tells the WaitGroup how many goroutines we expect to be working on processing the input channels.
	wg.Add(len(cs))

	// Start an output goroutine for each input channel in cs.
	// Output copies values from c to out until c is closed, then calls wg.Done.
	// output := func(c <-chan int) {
	// 	for n := range c {
	// 		out <- n
	// 	}
	// 	wg.Done()
	// }

	for _, c := range cs {
		go func(ch <-chan int) {
			for n := range ch {
				out <- n
			}

			// Inside each goroutine, after processing its respective input channel, it calls wg.Done() to decrement the WaitGroup, indicating that it has finished its work.
			wg.Done()
			// close(out) If you close out inside each goroutine processing an input channel, each goroutine will try to close the out channel independently when it finishes processing its respective input channel. This means that out might be closed multiple times concurrently, which can lead to a panic due to attempting to close a closed channel, or data races.

		}(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		// Meanwhile, another goroutine is waiting for all the goroutines to finish their work by calling wg.Wait(). This goroutine blocks until the WaitGroup counter becomes zero, which happens when all goroutines have called wg.Done().
		// After wg.Wait() returns, we can safely close out knowing that all input channels have been fully processed and no more data will be sent on out.
		wg.Wait()
		close(out)
	}()

	return out
}
