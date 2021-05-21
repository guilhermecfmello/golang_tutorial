package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		factorialPreview, _ := factorial(n - 1)
		return n * factorialPreview, nil
	}
}

func main() {
	result, _ := factorial(5)

	fmt.Println(result)
	_, err := factorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
