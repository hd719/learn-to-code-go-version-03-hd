package main

import "fmt"

func main() {
	in := gen() // creates a channel with int values

	c := factorial(in)

	for n := range c {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// for j := 3; j < 13; j++ {
			// 	out <- j
			// }
			out <- i
		}
		close(out)
	}()

	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()

	return out
}

func fact(n int) int {
	total := 1

	for i := n; i > 0; i-- {
		total = total * i
	}

	return total
}
