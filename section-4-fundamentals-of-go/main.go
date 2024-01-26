package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work because e is not used and go does not allow unused variables
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/

	// this does work
	/*
		var g int -> Declare a variable and assigns the zero value of type int to it
		fmt.Println(g)
		g = 43 -> Assign/Initialize a value to the variable
		fmt.Println(g)

		// declare a variable to hold a VALUE of a certain TYPE
		// initializing a variable
		var h int = 44
		fmt.Println(h)
	*/

	// ** Hexidecimal **
	// Difference between printf and println is that printf allows you to format the output of the string using verbs like %b and %x println allows you to print a line of text

	// adams := 42
	// fmt.Printf("42 as binary, %b \n", adams)
	// fmt.Printf("42 as hexadecimal, %x \n", adams)

	// a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	// fmt.Printf("%v \t %b \t %#x \n", a, a, a)
	// fmt.Printf("%v \t %b \t %#x \n", b, b, b)
	// fmt.Printf("%v \t %b \t %#x \n", c, c, c)
	// fmt.Printf("%v \t %b \t %#x \n", d, d, d)
	// fmt.Printf("%v \t %b \t %#x \n", e, e, e)
	// fmt.Printf("%v \t %b \t %#x \n", f, f, f)

	// ** Values, types, conversion, scope, & housekeeping **
	// y := 42
	// z := 42.0
	// fmt.Printf("%v of type %T \n", y, y)
	// fmt.Printf("%v of type %T \n", z, z)

	// var m float32 = 43.742
	// fmt.Printf("%v of type %T \n", m, m)

	/*
		// this does not work!
		// in go you can't take a VALUE that is float32 and store it
		// in a variable that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v of type %T \n", z, z)
	*/

	/*
		// this does work
		z = float64(m)
		fmt.Printf("%v of type %T \n", z, z)
	*/
}
