package main

import (
	"fmt"
	"test/bmw"
	"test/toyota"
)

type car interface {
	Process()
	Total() float64
}

func amount(cars ...car) float64 {
	var total float64
	for _, row := range cars {
		row.Process()
		total += row.Total()
	}

	return total
}

func main() {
	car1 := bmw.New(1000)
	car2 := toyota.New(1000)
	total := amount(car1, car2)
	fmt.Println(total)
}
