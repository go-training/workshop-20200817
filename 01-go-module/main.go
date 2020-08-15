package main

import (
	"fmt"
	"log"
	"os"

	hello "github.com/go-training/workshop-20200817/00-hello-module"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(hello.World())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	fmt.Println(s3Bucket)
}
