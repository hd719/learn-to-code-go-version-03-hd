package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Hamel"), boring("Desai"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // c is going block until it gets a value and until the loop runs 10 times
	}

	fmt.Println("You are both boring")
}

// This is not fan out because we do not channels coming into the function (however it is a pipeline)
func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ { // Syntax for infinite loop
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

// Fan in function to combine the two channels into one
func fanIn(input1 <-chan string, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1 // Get the value from input 1 and send it to c
		}
	}()

	go func() {
		for {
			c <- <-input2 // Get the value from input 2 and send it to c
		}
	}()

	return c
}
