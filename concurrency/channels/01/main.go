package main

import "fmt"

func main() {
	fmt.Println("Start Main method")

	ch := make(chan int)
	go myFunc(ch)
	ch <- 10

	fmt.Println("End Main method")
}

func myFunc(ch chan int) {
	fmt.Println(100 + <-ch)
}
