package main

import "fmt"

type item struct {
	productId int
	qtd       int
	price     float64
}

type order struct {
	userId int
	itens  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.itens {
		total += item.price * float64(item.qtd)
	}
	return total
}

func main() {
	order := order{
		userId: 1,
		itens: []item{
			item{1, 2, 3.98},
			item{2, 20, 10.25},
			item{3, 100, 99.0},
		},
	}

	fmt.Printf("Valor total: %.2f", order.totalValue())
}
