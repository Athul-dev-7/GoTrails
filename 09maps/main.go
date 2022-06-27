package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["JV"] = "Java"
	languages["SOL"] = "Solidity"
	fmt.Println("Languages are : ", languages)
	fmt.Println("PY is ", languages["PY"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
