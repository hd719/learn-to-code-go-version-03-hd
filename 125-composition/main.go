package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type foo int

// Composition is embedding a type into another type (similar to inheritance)
func main() {
	var a foo = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Printf("%T \n", p1)
	fmt.Printf("%T \n", p2)
	fmt.Printf("%#v \n", p2)

	p2 = p1 // This is not allowed because p1 is an anonymous struct and p2 is a person struct
	fmt.Printf("%T \n", p2)
	fmt.Printf("%#v \n", p2)
	// returns

}
