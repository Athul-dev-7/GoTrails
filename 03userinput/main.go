package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input section"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for my tutorials : ")

	// comma ok / comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating ,", input)
	fmt.Printf("Type of rating is a %T \n", input)
}
