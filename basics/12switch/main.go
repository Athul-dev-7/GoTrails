package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch cases in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(7)

	fmt.Println("Value of dice is ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Value is 1 and can start now")
	case 2:
		fmt.Println("Value is 2 and can move 2 spots")
	case 3:
		fmt.Println("Value is 3 and can move 3 spots")
	case 4:
		fmt.Println("Value is 4 and can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("Value is 5 and can move 5 spots")
		fallthrough
	case 6:
		fmt.Println("Value is 6, can move 6 spots and roll again")
	default:
		fmt.Println("What was that !")
	}
}
