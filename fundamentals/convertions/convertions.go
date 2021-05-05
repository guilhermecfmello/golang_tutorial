package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	x := 2.4
	y := 2

	// Converting int to float
	fmt.Println(x / float64(y))

	grade := 6.9
	finalGrade := int(grade)

	// danger
	fmt.Println("Test: " + string(49))

	// int to string
	fmt.Println("Test: " + strconv.Itoa(finalGrade))

	// string to int
	age := "123"
	// Atoi returns the result and a error (2 returns)
	num, _ := strconv.Atoi(age)
	fmt.Println("String ", age)
	fmt.Println("converted int type: ", reflect.TypeOf(num))
	fmt.Println("Int - 10: ", num-10)

	// string to boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("b is true")
	}

}
