package main

import "fmt"

type sportive interface {
	turnOnTurbo()
}

type luxuous interface {
	autoParking()
}

type sportiveLuxuous interface {
	sportive
	luxuous
	// Another methods
}

type bmw struct{}

func (b bmw) turnOnTurbo() {
	fmt.Println("Turning turbo on...")
}

func (b bmw) autoParking() {
	fmt.Println("Auto parking...")
}

func main() {
	var car sportiveLuxuous = bmw{}

	car.turnOnTurbo()
	car.autoParking()

}
