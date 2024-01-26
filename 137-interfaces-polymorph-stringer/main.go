package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c))) // convert the type count to the type int then to a string
}

type Stringer interface {
	String() string
}

func foo(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)

	foo(b)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
