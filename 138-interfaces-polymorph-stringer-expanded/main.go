package main

import (
	"fmt"
	"log"
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
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) { // takes a string type as a parameter that implements the String() method
	log.Println("LOG FROM 138", s.String())
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var n count = 42

	logInfo(b)
	logInfo(n)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
