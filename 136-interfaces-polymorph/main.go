package main

import "fmt"

type person struct {
	first string
}

// In this case secretAgent is a type and person is a type, they both have methods speak()
type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

// Interfaces allow us to have Polymorphism
type human interface {
	speak() // If you have this method(s), you are my type
}

func saySomething(h human) {
	h.speak()
}

func main() {

	// Since secretAgent type has a method speak(), it is also a type of human
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	// Since person type has a method speak(), it is also a type of human
	p2 := person{
		first: "Jenny",
	}

	sa1.speak() // This is a method that will return "I am a secret agent James"
	p2.speak()  // This is a method that will return "I am Jenny"

	// saySomething() is a function that takes in a human type (that are different values)
	saySomething(sa1)
	saySomething(p2)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { code }
