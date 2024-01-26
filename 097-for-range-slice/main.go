package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)

	// blank identifier to not use a returned value
	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}

	// same as above
	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}

	fmt.Println("-------------------")
	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])

	// doesn't work bc the index is out of range
	// fmt.Println(xs[3])

	fmt.Println(len(xs))
}
