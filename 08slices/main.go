package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Golang slices !")

	var fruitList = []string{"Apple", "Orange", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Tomato")
	fmt.Println("fruitList is ", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("fruitList is ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("fruitList is ", fruitList)

	fruitList = append(fruitList[:1])
	fmt.Println("fruitList is ", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 897
	highScores[1] = 947
	highScores[2] = 654
	highScores[3] = 325
	// highScores[4] = 832  //Can't add values like this
	fmt.Println("highScores are ", highScores)

	highScores = append(highScores, 652, 989, 313)
	fmt.Println("highScores are ", highScores)

	// sort slice
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted highScores are ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
