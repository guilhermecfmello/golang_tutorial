package main

import (
	"fmt"
	// Import "alias"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var radio = 3.2 // (float64) type decided by compiler

	// reduced way of var declaration
	area := PI * m.Pow(radio, 2)
	fmt.Println("Circunference area: ", area)

	// Several declarations
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	// reduced way of vars declarations
	g, h, i := 2, false, "Wat??"

	fmt.Println(g, h, i)

}
