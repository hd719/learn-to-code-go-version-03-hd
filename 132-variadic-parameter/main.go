package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", x)
}

// func (r receiver) identifier(p parameters) (return(s)) {code}

func sum(ii ...int) int { // Variadic parameter - accepts 0 or more values of a certain type (pass in any number of ints)
	fmt.Println(ii)
	fmt.Printf("%T\n", ii) // []int (slice of int)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
