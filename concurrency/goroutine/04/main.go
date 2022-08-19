package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	statements := []string{"Statement 1", "Statement 2", "Statement 3", "Statement 4", "Statement 5"}

	for _, stat := range statements {
		go displayStatements(stat)
	}

	time.Sleep(2 * time.Second)
}

func displayStatements(stat string) {
	fmt.Println(stat)
	time.Sleep(time.Second * 1)
}
