package main

import "fmt"

func main() {
	// create slice (internally is created an array)
	s := make([]int, 10)
	s[9] = 123
	fmt.Println(s)

	// create a slice with 10 elements but internally is created an array of 20 elements
	s = make([]int, 10, 20)

	fmt.Println(s, len(s), cap(s))

	// This code below doesnt work
	// fmt.Println("Printing slice overflowing index")
	// for i := 0; i < cap(s); i++ {
	// 	fmt.Println(s[i])
	// }

	s = append(s, 1, 2, 3, 4, 5, 6)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s, len(s), cap(s))
}
