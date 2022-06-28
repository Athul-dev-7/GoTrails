package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// simple for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("")

	// using range
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("")
	// for each in go
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// while loop in go
	fmt.Println("")
	value := 1
	for value < 10 {

		if value == 8 {
			goto num
		}
		if value == 5 {
			value++
			continue

		} else if value == 9 {
			fmt.Println("Value is 9, so breaking out from the loop")
			break
		} else {
			fmt.Println("Value is ", value)
			value++
		}

	}

num:
	fmt.Println("Hits by goto statement")

}
