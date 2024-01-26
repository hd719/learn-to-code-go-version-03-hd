package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, add) // pass the function add as an argument
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

// Function calls go at the end of parameter list
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
