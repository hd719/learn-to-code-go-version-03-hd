package main

import "fmt"

func main() {
	c1 := incrementor("Foo: ")
	c2 := incrementor("Bar: ")
	c3 := puller(c1)
	c4 := puller(c2)

	fmt.Println("Final counter", <-c3+<-c4) // Blocks main go routine until c3 and c4 are closed
}

func incrementor(s string) chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			out <- i
			fmt.Println(s, i)
		}
		close(out)
	}()

	return out
}

func puller(c chan int) chan int {
	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()

	return out
}
