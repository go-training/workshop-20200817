package main

import (
	"fmt"
	"time"
)

func main() {
	example01()
	example02()
	example03()
}

// buffer vs unbuffer channel
func example03() {
	c := make(chan bool)
	go func() {
		fmt.Println("send to channel")
		c <- true
	}()
	<-c
	fmt.Println("Done")
}

func example01() {
	c := make(chan string)
	go foobar("example01", c)
	for {
		v, ok := <-c
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

func example02() {
	c := make(chan string)
	go foobar("example02", c)
	for v := range c {
		fmt.Println(v)
	}
}

func foobar(v string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- v
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}
