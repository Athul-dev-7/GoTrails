package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev"

func main() {
	fmt.Println("Handling Web Request in Golang")

	response, err := http.Get(URL)
	checkNilErr(err)
	fmt.Printf("Response is of type %T\n", response) // type will be a pointer
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	content := string(databytes)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
