package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang arrays !")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Peach"
	fruitList[3] = "Orange"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Length of fruitList is ", len(fruitList))

	var vegList = [4]string{"Beans", "Potato", "Mushroom"}
	fmt.Println("Veg list is ", vegList)
	fmt.Println("Length of veg list is ", len(vegList))

}
