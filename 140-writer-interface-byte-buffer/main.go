package main

import (
	"bytes"
	"fmt"
)

// type person struct {
// 	first string
// }

// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer f.Close()

	// s := []byte("Hello gophers!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())
	b.WriteString("gophers!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("Its a new string")
	fmt.Println(b.String())
	b.Reset()
	b.Write([]byte("Hello gophers!"))
	fmt.Println(b.String())
}
