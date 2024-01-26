package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...) // Unfurling a slice, we are passing in the values of a slice instead of the entire slice
	// xi... is the same as 1, 2, 3, 4, 5, 6, 7, 8, 9
	// You would use this if you already have a slice and you want to pass in the values of the slice to a variadic parameter
	fmt.Println("The sum is", x)
}

// func (r receiver) identifier(p parameters) (return(s)) {code}

func sum(ii ...int) int {
	fmt.Println("ii is", ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
