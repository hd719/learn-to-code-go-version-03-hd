package main

import "fmt"

func main() {
	// defer bazz() returns // bar foo bazz
	defer foo() // defer foo() until the end of the enclosing function (main)
	// Runs after everything else in the enclosing function (main) has finished
	defer bazz() // returns bar bazz foo

	// Order of defers is last in first out
	bar()
}

// func (r receiver) identifier(p parameter(s)) (r return(s)) { code }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func bazz() {
	fmt.Println("bazz")
}
