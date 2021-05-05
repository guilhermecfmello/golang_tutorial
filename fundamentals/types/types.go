package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Integers
	fmt.Println(1, 2, 1000)

	// reflect.Typeof: Get the variable type
	fmt.Println("Literal integer: ", reflect.TypeOf(32000))

	// Sinal less (positives only)... uint8 uint16 uint32 uint64
	// byte = alias for int
	var b byte = 255
	fmt.Println("Byte: ", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("Max value of int: ", i1)
	fmt.Println("i1 type: ", reflect.TypeOf(i1))

	// unicode table representation
	var i2 rune = 'b'
	fmt.Println("i2: ", i2)
	fmt.Println("Rune: ", reflect.TypeOf(i2))

	// floats (float32, float34)
	var x float32 = 49.99
	// or
	// var x = float32(49.99)
	fmt.Println("x type: ", reflect.TypeOf(x))
	fmt.Println("Type of literal 49.99: ", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("bo type: ", bo)
	fmt.Println(!bo)

	// string is "", not ''
	s1 := "Hello World!"
	fmt.Println(s1)
	fmt.Println("String size: ", len(s1))
	// string with break lines
	s2 := `My
	name is
	Jonasss
	sss`
	fmt.Println(s2)

	// char '', but is int32 remaped
	char := 'a'
	fmt.Println("Type of char: ", reflect.TypeOf(char))
	fmt.Println("char: ", char)
	fmt.Printf("char formated as (char): %c", char)
}
