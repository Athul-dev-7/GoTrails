package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Statement 1")

	go func() {
		fmt.Println("Statement 2")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Statement 3")
}
