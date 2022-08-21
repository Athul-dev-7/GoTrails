package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("---Main Goroutine Starts---")

	go authorName()
	go authorIds()

	time.Sleep(time.Millisecond * 3500)
	fmt.Println("---Main Goroutine Ends---")
}

func authorName() {
	names := [5]string{"Rohit", "Suman", "Aman", "Ria"}
	for i := 0; i < len(names); i++ {
		time.Sleep(time.Millisecond * 150)
		fmt.Printf("%s\n", names[i])
	}
}

func authorIds() {
	ids := [4]int{100, 101, 102, 103}
	for i := 0; i < len(ids); i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("%d\n", ids[i])
	}
}
