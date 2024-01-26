package main

import "fmt"

func main() {
	x()

	y := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	s := func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	s()
	y()
}

var x = func() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("foo ran")
}
