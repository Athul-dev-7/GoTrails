package main

import "fmt"

func main() {

	ch := make(chan string)

	go func() {
		ch <- "A"
		ch <- "B"
		ch <- "C"
		ch <- "D"
		ch <- "E"
		close(ch)
	}()

	for item := range ch {
		fmt.Println(item)
	}
}
