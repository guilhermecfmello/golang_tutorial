package main

import "fmt"

type course struct {
	name string
}

func main() {
	var c interface{}
	fmt.Println(c)

	c = 3
	fmt.Println(c)

	type dinamic interface{}

	var c2 dinamic = "asd"
	fmt.Println(c2)

	c2 = true
	fmt.Println(c2)

	c2 = "New value for c2"
	fmt.Println(c2)
}
