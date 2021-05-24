package main

import "fmt"

// type <NAME_STRUCT> struct
type product struct {
	name     string
	price    float64
	discount float64
	status   bool
}

// Receiver
func (p product) priceWithDiscount() float64 {
	return p.price * (1 - (p.discount / 100))
}

func main() {
	var product1 product
	product1 = product{
		name:     "Lapis",
		price:    1.98,
		discount: 10,
		status:   true,
	}

	fmt.Println(product1, product1.priceWithDiscount())

	product2 := product{"Notebook", 1989.90, 0.10, true}
	fmt.Println(product2)
}
