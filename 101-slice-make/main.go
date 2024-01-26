package main

import "fmt"

func main() {
	// si := []string{"a", "b", "c"} // slice literal
	// fmt.Println(si)

	// make a slice instead of using the literal syntax
	// 1st arg: type of slice
	// 2nd arg: length of the slice
	// 3rd arg: how many elements to allocate in the slice
	// For example, make([]int, 0, 10) allocates an underlying array
	// of size 10 and returns a slice of length 0 and capacity 10 that is
	// backed by this underlying array.
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println("------------")
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("------------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}

/*
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("-------------")
*/
