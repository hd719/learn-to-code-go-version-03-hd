package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Print("a: ", *a, "\n")
	fmt.Print("b: ", *b, "\n")
	fmt.Print("c: ", *c, "\n")
	fmt.Print("d: ", *d, "\n")
}
