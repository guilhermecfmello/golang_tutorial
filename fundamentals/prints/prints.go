package main

import "fmt"

func main() {

	x := 3.141516

	var xs string = fmt.Sprint(x)

	// Cacatenating strings with "+" operation
	fmt.Println("X value: " + xs)
	// Printings x float as argument
	fmt.Println("X value:", x)

	fmt.Printf("X value formated: %.2f\n", x)
	fmt.Printf("X value formated as string: %s\n", xs)

	a := 1
	b := 1.999
	c := false
	d := "oops"
	// 					int		float		float_formated		boolean		string
	fmt.Printf("%d 		%f 					%.1f 						%t 				%s		\n", a, b, b, c, d)
	// Generic print
	fmt.Printf("%v %v %v %v", a, b, c, d)
}
