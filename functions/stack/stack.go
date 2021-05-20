package main

import (
	"fmt"
	"runtime/debug"
)

func f3() {
	fmt.Println("Starting F3")
	debug.PrintStack()
	fmt.Println("Ending F3")
}

func f2() {
	fmt.Println("Starting F2")
	f3()
	fmt.Println("Ending F2")
}

func f1() {
	fmt.Println("Starting F1")
	f2()
	fmt.Println("Ending F1")
}

func f4(n int) int {
	if n < 1 {
		return 0
	}
	fmt.Println("N: ", n)
	return f4(n - 1)
}

func main() {
	f1()

	f4(10)
}
