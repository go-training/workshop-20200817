package main

import (
	"fmt"
	"sync"
	"time"
)

func foo(input string) {
	fmt.Println(input)
}

func main() {
	// one
	go foo("1")
	go foo("2")

	// two
	msg := "let's go"
	go func() {
		fmt.Println(msg)
	}()
	msg = "let's gogogo"

	go waitGroup()
	time.Sleep(2 * time.Second)
}

func waitGroup() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("m:", len(m))
	fmt.Println("m:", m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
