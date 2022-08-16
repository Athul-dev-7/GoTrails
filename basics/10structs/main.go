package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance in golang; no super or parent

	athul := User{"Athul", "athul@gmail.com", true, 16}
	fmt.Println(athul)

	fmt.Println("")
	var pers1 Person
	var pers2 Person

	pers1.name = "Abin"
	pers1.age = 22
	pers1.job = "Frontend Developer"
	pers1.salary = 35000
	// fmt.Println("Person1 is :", pers1)
	printPerson(pers1)

	fmt.Println("")

	pers2.name = "Jithin"
	pers2.age = 24
	pers2.job = "Backend Developer"
	pers2.salary = 5500
	// fmt.Println("Person1 is :", pers2)
	printPerson(pers2)
}

func printPerson(pers Person) {
	fmt.Println("Name :", pers.name)
	fmt.Println("Age :", pers.age)
	fmt.Println("Job :", pers.job)
	fmt.Println("Salary :", pers.salary)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Person struct {
	name   string
	age    int
	job    string
	salary int
}
