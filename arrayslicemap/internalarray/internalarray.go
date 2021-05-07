package main

import "fmt"

// Prooving that the internal array is the same to different slices
func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	fmt.Println(s1, s2)

	s1[0] = -1
	fmt.Println(s1, s2)
	fmt.Println(s1, len(s1), cap(s1))

	s3 := append(s2, 1, 2, 3, 4, 5, 6, 7, 8)
	s3[2] = 999
	s2[2] = 888
	// same internal array
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// reallocated internal array
	fmt.Println(s3, len(s3), cap(s3))
}
