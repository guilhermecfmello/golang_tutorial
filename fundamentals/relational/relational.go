package main

import (
	"fmt"
	"time"
)

func main() {
	var x int = 2
	var y int = 3

	fmt.Println("Strings: ", "Gui" == "Gui2")
	fmt.Println("x: ", x, "y: ", y)
	fmt.Println("!=", x != y)
	fmt.Println("<", x < y)
	fmt.Println(">", x > y)
	fmt.Println("<=", x <= y)
	fmt.Println(">=", x <= y)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas: ", d1 == d2)
	fmt.Println("Datas: ", d1.Equal(d2))

	type People struct {
		Name string
	}

	p1 := People{"Gui"}
	p2 := People{"Gui"}

	fmt.Println("People: ", p1 == p2)
}
