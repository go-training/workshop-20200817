package main

import (
	"fmt"
	"time"
)

func main() {
	go foobar("foobar")
	time.Sleep(2 * time.Second)
}

func foobar(v string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(v)
		time.Sleep(500 * time.Millisecond)
	}
}
