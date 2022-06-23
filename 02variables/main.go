package main

import "fmt"

// constants or public var
const LoginToken = "akjrbgqjergoquirgboqlrgbqo34rgoikjn"

func main() {
	fmt.Println("Variables")

	var username string = "athul"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n", smallValue)

	var smallFloat1 float32 = 255.45124585498756464146546
	fmt.Println(smallFloat1)
	fmt.Printf("Variable is of type : %T \n", smallFloat1)

	var smallFloat2 float64 = 255.45124585498756464146546
	fmt.Println(smallFloat2)
	fmt.Printf("Variable is of type : %T \n", smallFloat2)

	// default values

	var anyNumber int
	fmt.Println(anyNumber)
	fmt.Printf("Variable is of type : %T \n", anyNumber)

	var anyBool bool
	fmt.Println(anyBool)
	fmt.Printf("Variable is of type : %T \n", anyBool)

	var anyString string
	fmt.Println(anyString)
	fmt.Printf("Variable is of type : %T \n", anyString)

	var anyFloat32 float32
	fmt.Println(anyFloat32)
	fmt.Printf("Variable is of type : %T \n", anyFloat32)

	var anyFloat64 float64
	fmt.Println(anyFloat64)
	fmt.Printf("Variable is of type : %T \n", anyFloat64)

	// implicit type
	var website = "www.xyz.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	// no var type (can only be used inside any functions not outside)
	numberOfUsers := 6500
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type : %T \n", numberOfUsers)

	// accessing Public var
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
