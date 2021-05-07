package main

import (
	"fmt"
	"reflect"
)

// Slice of an array, we use it as a pointer to an array or an especific element's array

func main() {
	a1 := [3]int{1, 2, 3} // array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, "type: ", reflect.TypeOf(a1), "||", s1, " type: ", reflect.TypeOf(s1))

	a2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s2 := a2[1:3]
	s2[0] = 123
	fmt.Println(a2, s2)

	s3 := a2[:4]

	s4 := a2[8:]

	fmt.Println(s3)
	fmt.Println(s4)

	s5 := a2[0]

	fmt.Println(s5)
}
