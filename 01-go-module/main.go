package main

import (
	"fmt"

	"hello/foo"

	hello "github.com/go-training/workshop-20200817/00-hello-module"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(hello.World())
	fmt.Println(foo.Hello())
}
