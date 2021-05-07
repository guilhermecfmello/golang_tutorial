package main

import "fmt"

func main() {

	x := 1
	y := 2

	// Go is only posfixed
	x++
	y = y + 1
	y--
	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)

	// Is denied to decrement in an expression
	// fmt.Println(x == y--)

	// Precedence level in Go has only 5 levels, while C has more than 13
}
