package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a*a == b {
		fmt.Print(a, " is the square of ", b)
	} else if b*b == a {
		fmt.Print(b, " is the square of ", a)
	} else {
		fmt.Print(a, " and ", b, " are not squares of each other")

	}
}
