package main

import (
	"fmt"
	"time"
)

func main() {

	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	select {
	case op1 := <-R1:
		fmt.Println(op1)

	case op2 := <-R2:
		fmt.Println(op2)

	}
}

func portal1(ch1 chan string) {
	time.Sleep(time.Second * 3)
	ch1 <- "Welcome to Portal 1"
}

func portal2(ch2 chan string) {
	time.Sleep(time.Second * 9)
	ch2 <- "Welcome to Portal 2"

}
