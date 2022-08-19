package main

import "fmt"

func main() {
	// Calling a Goroutine
	go display("welcome")

	// Calling a Normal function
	display("Hello, World")
}

func display(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
}
