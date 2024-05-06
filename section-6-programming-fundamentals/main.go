package main

// the imported package "fmt"
// is in the "file block" scope
// of this file
import (
	"fmt"

	"example.com/furtherexplored"
)

// x is in "package block" scope
var x = 42

func main() {
	// x can be accessed here
	fmt.Println(x)

	// x can be accessed here
	printMe()

	// y does not exist here yet
	// so this won't work
	// fmt.Println(y)

	// y is in "block scope"
	y := 43
	fmt.Println(y)

	p1 := person{
		"James",
	}
	p1.sayHello()

	// variable "shadowing" x is defined in package level scope and block level scope
	x := 32
	fmt.Println("PRINTING", x) // outputs 32

	// package block x is still the same
	printMe() // outputs 42

	furtherexplored.Fascinating() // being called from another package

	//also good to know

	if z := 82; z > 50 {
		fmt.Println(z)
	}
	// you can't access z here
	// fmt.Println(z)
	/*
		take a look at the "comma ok idiom"
		https://go.dev/doc/effective_go#maps
	*/
}

func printMe() {
	//x can be accessed here
	fmt.Println(x)
}

// type person is in "package block" scope
type person struct {
	first string
}

// the method sayHello
// which is attached to VALUES of TYPE person
// is in "package block" scope // outputs "Hi, my name is James"
func (p person) sayHello() {
	fmt.Println("Hi, my name is", p.first)
}
