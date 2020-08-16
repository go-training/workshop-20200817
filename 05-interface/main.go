package main

import (
	"errors"
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

	ips := map[string]ipv4{
		"google":    {1, 1, 1, 1},
		"localhost": {127, 0, 0, 1},
	}

	for s, v := range ips {
		fmt.Printf("%v, %v\n", s, v)
	}

	printF(1)
	printF("1")
}

type ipv4 []byte

func (s ipv4) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", s[0], s[1], s[2], s[3])
}

// type assertion
func printF(val interface{}) error {
	if v, ok := val.(string); ok {
		fmt.Println("this is string via if:", v)
	} else if v, ok := val.(int); ok {
		fmt.Println("this is int via if:", v)
	}

	switch v := val.(type) {
	case string:
		fmt.Println("this is string via switch:", v)
	case int:
		fmt.Println("this is int via switch:", v)
	}

	return errors.New("Error: this is not string")
}
