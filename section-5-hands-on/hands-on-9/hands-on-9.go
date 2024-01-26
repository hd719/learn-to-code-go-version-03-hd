package hands_on_9

import "fmt"

func handsOn9() {
	fmt.Println("Hello, playground")

	const (
		_  = iota
		KB = 1 << (iota * 10)
		MB
		GB
		TB
		PB
	)

	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
}
