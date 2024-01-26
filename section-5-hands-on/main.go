package main

import (
	"fmt"
	"math"
)

// Bitwise Ops (Bitshifs)
// Numeric constants are high-precision values.
// An untyped constant takes the type needed by its context.
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	// 1 << 100 = 2^100 = 1267650600228229401496703205376
	BigNumber = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	SmallNumber = BigNumber >> 99
)

// iota is a special predeclared identifier that can be used in a constant declaration to simplify definitions of incrementing numbers
// iota starts at 0 in each const block and increments by one
const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
	c3 = iota // c3 == 3
)

const (
	_ = iota // a == 0
	a
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(math.Pi) // A name is exported if it begins with a capital letter, eq to import/export in JS or TS]

	// Below the left column shows us binary numbers, and the right column shows us the decimal representation of those numbers in base 10
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)

	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)

}

// The super (possibly over) simplified definition is just that << is used for "times 2" and >> is for "divided by 2" - and the number after it is how many times.

// So n << x is "n times 2, x times". And y >> z is "y divided by 2, z times".

// For example, 1 << 5 is "1 times 2, 5 times" or 32. And 32 >> 5 is "32 divided by 2, 5 times" or 1.
