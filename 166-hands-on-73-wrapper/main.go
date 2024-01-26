package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func doWork() {
	for i := 0; i < 2000; i++ {
		fmt.Println(i)
	}
}

type funcType func()

// func timeFunc(f func()) {
func timeFunc(f funcType) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
