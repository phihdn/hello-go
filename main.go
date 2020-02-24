package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Print from the Go program")
	testEnv := os.Getenv("TEST_ENV")
	if testEnv == "" {
		testEnv = "World"
	}
	fmt.Println("Hello", testEnv)
}
