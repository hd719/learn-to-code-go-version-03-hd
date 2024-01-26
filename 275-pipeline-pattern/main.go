package main

import "fmt"

func main() {
	// Set up pipeline

	// c := gen(2, 3)
	// out := sq(gen(2, 3))

	// Consume the channel
	// fmt.Println(<-out) // 4
	// fmt.Println(<-out) // 9

	for n := range sq(sq(sq(gen(2, 3)))) {
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

func sq(in chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}