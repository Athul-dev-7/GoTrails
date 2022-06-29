package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")

	//create a file sample.txt using os
	content := "This is a sample file"
	file, err := os.Create("./sample.txt")
	checkNilErr(err)

	// write content to the sample.txt using io
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length of file is :", length)
	defer file.Close()
	readFile("./sample.txt")

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text inside sample.txt is :", string(dataByte))
}
