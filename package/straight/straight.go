package main

import "math"

// Inicializing with uppercase (variables, functions, constants) = public => Variable
// Inicializing with lowercase (variables, functions, constants) = private => variable or _Variable

type Point struct {
	x float64
	y float64
}

func catetos(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Linear Distance between two points
func Distance(a, b Point) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
