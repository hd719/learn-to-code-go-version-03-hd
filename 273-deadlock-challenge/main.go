package main

import "fmt"

// Challenge 1: Fix the deadlock
// func main() {
// 	c := make(chan int)

// 	// c <- 1 // Blocks main go routine

// 	go func() {
// 		c <- 1
// 	}()

// 	fmt.Println(<-c)
// }

// Challenge 2: Fix the deadlock
// func main() {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()

// 	for v := range c {
// 		fmt.Println(v)
// 	}

// }

// Challenge 3: Fix the deadlock
func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

// func factorial(n int) int {
// 	total := 1

// 	for i := n; i > 0; i-- {
// 		total *= i
// 	}

// 	return total
// }

func factorial(n int) int {
	out := make(chan int)
	total := 1

	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()

	for v := range out {
		total *= v
	}

	return total
}
