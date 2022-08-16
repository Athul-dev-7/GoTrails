package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Golang time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("Fomatted time is ", formattedTime)

	createdDate := time.Date(2022, time.November, 7, 15, 25, 35, 0, time.UTC)
	fmt.Println("Created date is ", createdDate.Format("01-02-2006 15:04:05 Monday"))
}
