package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers !")

	var one int
	fmt.Println("Value of variable one is ", one)

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	var myNumber int = 14
	var ptr = &myNumber
	fmt.Println("Value of myNumber is ", myNumber)
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of pointer is ", *ptr)

	*ptr = *ptr + 6
	fmt.Println("New value is ", *ptr)
	fmt.Println("New value is ", myNumber)

}
