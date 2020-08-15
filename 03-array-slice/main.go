package main

import "fmt"

func main() {
	foo01()
	foo02()
	foo03()
	foo04()
}

func foo01() {
	var foo [3]int
	var bar []int

	if bar == nil {
		fmt.Printf("%p\n", foo)
		fmt.Printf("%p\n", bar)
	}
}

func foo02() {
	x := []int{1, 2, 3}
	y := x[1:]
	y[0] = 100
	fmt.Println(x)
	fmt.Println(y)
	y = append(y, 1000)
	y[0] = 99
	fmt.Println(x)
	fmt.Println(y)
}

func foo03() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
}

func foo04() {
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(x)
	}(x)
	fmt.Println(x)
}
