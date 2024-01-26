package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// ~ is used to define underlying types
type myNumbers interface {
	~int | ~float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {
	// var n myAlias = "Test string" // Doesnt work doesnt satisfy the type constraint
	var n myAlias = 42

	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	// Without the underlying type constraint (~), the following line would not work
	// myAlias does not satisfy myNumbers (possibly missing ~ for int in myNumbers)
	// In below line, n is of type myAlias (which is an int) and the add function expects a type that satisfies the myNumbers interface which is int | float64
	// By adding the ~ to int in the myNumbers interface, in the interface of myNumbers you will accept any type that is an int or any type that is an underlying type of int
	// fmt.Println(addT(n, 2))
	fmt.Println(addT(n, 2))
	fmt.Println(addT(1.2, 2.2))

}
