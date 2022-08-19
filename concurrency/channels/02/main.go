package main

import "fmt"

func main() {

	ch := make(chan string)
	go myFunc(ch)

	for {
		res, ok := <-ch
		if !ok {
			fmt.Println("Channel  CLOSED")
			break
		}
		fmt.Println("Channel OPEN", res, ok)

	}

}

func myFunc(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "Hello"
	}
	close(ch)
}
