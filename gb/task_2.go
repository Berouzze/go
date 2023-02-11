package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a < b {
		fmt.Print(a, " меньше ", b)
	} else if b == a {
		fmt.Print(b, " и ", a, " is the square of ")

	} else {
		fmt.Println(b, " меньше ", a)
	}
}
