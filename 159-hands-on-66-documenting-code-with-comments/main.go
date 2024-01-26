package main

import "fmt"

func main() {
	fmt.Println(Paradise("Hawaii"))
}

// Paradise prints out the user's input of paradise as a sentence.
// go doc -cmd Paradise
// go doc -src Paradise
// go doc -cmd -all
func Paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}
