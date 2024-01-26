package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi - %v\n", xi)  // %v - print the value in a default formats
	fmt.Printf("xi - %#v\n", xi) //#v - print the value in a Go-syntax representation
	fmt.Println("-------------")

	// [inclusive:exclusive]
	fmt.Printf("xi - %#v\n", xi[0:4]) // inlcude 0 and exclude 4 (index)
	fmt.Println("-------------")

	// [:exclusive]
	fmt.Printf("xi - %#v\n", xi[:7]) // inlcude 0 and exclude 7 (index)
	fmt.Println("-------------")

	// [inclusive:]
	fmt.Printf("xi - %#v\n", xi[4:]) // inlcude 4 and exclude the last element (index)
	fmt.Println("-------------")

	// [:]
	fmt.Printf("xi - %#v\n", xi[:]) // inlcude 0 and exclude the last element (index)
	fmt.Println("-------------")
}
