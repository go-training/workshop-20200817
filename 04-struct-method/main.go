package main

import (
	"fmt"
	"log"
	"time"

	"test/car"
	"test/email"
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

	e, err := email.New("foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	if err := e.Send("email@gmail.com"); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
}
