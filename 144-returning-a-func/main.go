package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x)

	y := bar(2)
	z := y()
	fmt.Println("Chained", bar(2)())
	fmt.Println(z)

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", y)
}

func foo() int {
	return 42
}

func bar(h int) func() int {
	return func() int {
		return h * 43
	}
}
