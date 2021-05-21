package main

import "fmt"

func getApprovedValue(number int) int {
	number2 := 0
	defer fmt.Println("end! Value is ", number2)
	if number > 5000 {
		fmt.Println("High value")
		number2 = 123
		return 5000
	} else {
		fmt.Println("Low value")
		number2 = 321
		return number
	}
}

func main() {
	fmt.Println(getApprovedValue(3000))
}
