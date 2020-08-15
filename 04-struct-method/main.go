package main

import (
	"fmt"
	"time"

	"test/car"
)

func main() {
	car1 := car.Car{
		Name:  "lexus",
		Color: "white",
		Price: 800,
	}

	fmt.Println("car1", car1.Name)
	fmt.Println("car1", car1.Color)
	car1.UpdateColor("yellow")
	fmt.Println("car1", car1.Color)

	car2 := car.New("test", "red")

	fmt.Println("car2", car2.Name)
	fmt.Println("car2", car2.Color)
	car2.UpdateColor("yellow")
	fmt.Println("car2", car2.Color)

	for i := 0; i < 10; i++ {
		go func(i int) {
			car2.SetPrice((i + 1) * 100)
			time.Sleep(23 * time.Millisecond)
			fmt.Println(car2.Price)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
