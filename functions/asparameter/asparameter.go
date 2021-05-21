package main

import "fmt"

func mul(a int, b int) int {
	return a * b
}

func exec(function func(int, int) int, p1, p2 int) int {
	return function(p1, p2)
}

func main() {
	a, b := 5, 7
	fmt.Println("Exec mul: ", exec(mul, a, b))

	sum := func(a, b int) int {
		return a + b
	}

	fmt.Println("Exec sum: ", exec(sum, a, b))
}
