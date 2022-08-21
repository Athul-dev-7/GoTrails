// In select statement, if multiple cases are ready to proceed, then one of them can be selected randomly
package main

import "fmt"

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
	for i := 0; i < 3; i++ {
		ch1 <- "Welcome to portal 1"
	}
}

func portal2(ch2 chan string) {
	ch2 <- "Welcome to portal 2"

}
