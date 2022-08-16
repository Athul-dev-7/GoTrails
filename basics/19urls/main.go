package main

import (
	"fmt"
	"net/url"
)

const URL = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=lajfhrgiourheg214356"

func main() {
	fmt.Println("Handling urls in Golang")
	fmt.Println(URL)

	// Extracting a url
	response, err := url.Parse(URL)
	checkNilErr(err)
	fmt.Println(response.Scheme)
	fmt.Println(response.Host)
	fmt.Println(response.Path)
	fmt.Println(response.Port())
	fmt.Println(response.RawQuery)

	params := response.Query()
	fmt.Printf("params is of type : %T\n", params)
	fmt.Println(params)

	for _, val := range params {
		fmt.Println("Param is ", val)
	}

	// constructing a url
	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/learn",
	}

	finalURL := partsOfURL.String()
	fmt.Println(finalURL)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
