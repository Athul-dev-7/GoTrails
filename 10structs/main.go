package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// No inheritance in golang; no super or parent

	athul := User{"Athul", "athul@gmail.com", true, 16}
	fmt.Println(athul)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
