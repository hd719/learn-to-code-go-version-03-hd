package main

import "fmt"

func main() {
	foo()

	bar("Todd") // Passing in an argument aka receiver

	sampleFunc()

	s := aloha("Todd")
	fmt.Println(s)

	s1, n := dogYears("Todd", 40) // Can return multiple values
	fmt.Println(s1, n)
}

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func sampleFunc() {
	fmt.Println("I am from sampleFunc")
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	// Sprint returns a string instead of printing it
	// Sprintf formats according to a format specifier and returns the resulting string.
	// Sprint formats using the default formats for its operands and returns the resulting string. Spaces are added between operands when neither is a string.
	return fmt.Sprint(name, " is this old in dog years "), age
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

/*

func syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params, 2 returns
*/
