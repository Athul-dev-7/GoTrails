package main

import "fmt"

func main() {
	fmt.Println("Defer in Golang")
	fmt.Println("")

	// default way
	fmt.Println("One")
	fmt.Println("Two")
	fmt.Println("Three")

	fmt.Println("")
	// when using defer
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

}
