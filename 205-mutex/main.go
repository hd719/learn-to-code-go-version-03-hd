package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Creating 100 go routines and incrementing the counter by 1, however this is creating a race concidtion because its all updating the same variable!
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex // Mutex is a lock that is used to lock a variable so only one go routine can access it at a time

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // Locks the variable so only one go routine can access it at a time
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched() // yields the processor, allowing other goroutines to run
			v++
			counter = v
			mu.Unlock() // Locks the variable so only one go routine can access it at a time
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
