package main

import "fmt"

func main() {
	var Number, lastNumber int
	fmt.Scan(&Number)
	lastNumber = Number % 10
	fmt.Println(lastNumber)
}
