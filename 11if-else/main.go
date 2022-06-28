package main

import "fmt"

func main() {
	fmt.Println("if else in Golang")

	inningsTotal := 248
	var result string

	if inningsTotal > 250 {
		result = "Score is defendable"
	} else if inningsTotal < 250 {
		result = "Score is below par"
	} else {
		result = "Something else"
	}
	fmt.Println(result)

	num := 9
	if num%2 == 0 {
		fmt.Println("Num is even")
	} else {
		fmt.Println("Num is not even")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is less than 10")
	}
}
