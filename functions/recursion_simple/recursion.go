package main

import "fmt"

func factorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		factorialPreview := factorial(n - 1)
		return n * factorialPreview
	}
}

func main() {
	result := factorial(5)

	fmt.Println(result)
	result = factorial(-2)
	fmt.Println(result)
}
