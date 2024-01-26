package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println("----------------")

	y := &x // Assign the memory address of x to y
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)  // Value of 42. Dereference the memory address of y to get the value of x
	fmt.Println(*&x) // Value of 42

	fmt.Println("----------------")

	*y = 43 // Basically y is the memory address of x. So, we are declaring take the value at the memory address of y and assign it to 43
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*y)
}
