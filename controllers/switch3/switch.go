package main

import (
	"fmt"
	"time"
)

// Function can receives an interface type as argument, a generic one
func tipe(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "function"
	default:
		return "Can't detect the variable type"
	}
}

func main() {

	fmt.Println("Type: " + tipe(2.3))
	fmt.Println("Type: " + tipe(4))
	fmt.Println("Type: " + tipe("Gui"))
	fmt.Println("Type: " + tipe(func() {}))
	fmt.Println("Type: " + tipe(time.Now()))
}
