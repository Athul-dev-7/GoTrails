package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is : ", result)

	proResult, proMessage := proAdder(2, 5, 7, 8, 9, 4, 5)
	fmt.Println("Pro result is :", proResult)
	fmt.Println("Pro message is :", proMessage)

}

func greeter() {
	fmt.Println("Good Morning")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi proAdder function"
}
