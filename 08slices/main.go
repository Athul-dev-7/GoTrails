package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Golang slices !")

	// 	There are several ways to create a slice:

	// *	Using the []datatype{values} format
	// *	Create a slice from an array
	// *	Using the make() function

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

	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "ruby", "python"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	// 	In Go, there are two functions that can be used to return the length and capacity of a slice:

	// *	len() function - returns the length of the slice (the number of elements in the slice)
	// *	cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	my_slice := []string{"Go", "Python", "Javascript", "Ruby", "Haskel"}
	fmt.Println("my_slice is :", my_slice)
	fmt.Println("Length of my_slice is :", len(my_slice))
	fmt.Println("Capacity of my_slice is :", cap(my_slice))

	// You can create a slice by slicing an array
	my_array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("my_array is :", my_array)

	myarr_slice := my_array[1:3]
	fmt.Println("myarr_slice is :", myarr_slice)
	fmt.Println("Length of my_array is :", len(my_array))
	fmt.Println("Length of myarr_slice is :", len(myarr_slice))
	fmt.Println("Capacity of myarr_slice is :", cap(myarr_slice))

	// The make() function can also be used to create a slice.

	slice := make([]int, 5, 8)
	fmt.Println("Slice is :", slice)
	fmt.Println("Length of slice is :", len(slice))
	fmt.Println("Capacity of slice is :", cap(slice))
}
