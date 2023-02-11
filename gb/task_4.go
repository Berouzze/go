package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if b < a && c < a {
		fmt.Print(a)
	} else if a < b && c < b {
		fmt.Print(b)
	} else {
		fmt.Print(c)

	}
}
