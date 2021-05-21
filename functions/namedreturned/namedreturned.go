package main

import "fmt"

func swap(p1, p2 int) (second int, first int) {
	second = p2
	first = p1
	return
}

func main() {
	v1 := 999
	v2 := 123

	v1, v2 = swap(v1, v2)
	fmt.Println("v1:", v1, "v2: ", v2)

	v1 = 9
	v2 = 10
	second, first := swap(v1, v2)
	fmt.Println("first:", first, "second: ", second)
}
