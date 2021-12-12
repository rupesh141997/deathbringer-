package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world ")
	fmt.Println("java ")

	fmt.Println("Sum:", sum(5, 6))
}

func sum(a int, b int) int {
	return a + b
}
