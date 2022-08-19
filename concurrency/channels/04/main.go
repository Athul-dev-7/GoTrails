package main

import "fmt"

func main() {
	ch := make(chan string, 5)

	ch <- "A"
	ch <- "B"
	ch <- "C"
	ch <- "D"
	ch <- "E"

	// To find the length of the channel
	fmt.Println(len(ch))

	// to find the capacity of the channel
	fmt.Println(cap(ch))
}
