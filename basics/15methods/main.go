package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")

	athul := User{"Athul", "athul@gmail.com", true, 25}
	fmt.Println("User athul is :", athul)
	fmt.Printf("User name is %v\n", athul.Name)
	fmt.Printf("User email is %v\n", athul.Email)
	fmt.Printf("User age is %v\n", athul.Age)
	fmt.Printf("User status is %v\n", athul.Status)
	athul.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}
