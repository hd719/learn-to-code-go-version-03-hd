package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := &person{
		first: "James",
	}

	p2 := person{
		first: "Jenny",
	}

	saySomething(p1)
	// saySomething(p2) // This will not work because the speak method is a pointer receiver
}
