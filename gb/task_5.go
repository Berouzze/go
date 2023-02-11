package main

import "fmt"

func main() {
	var a, c, b int
	fmt.Scan(&a, &c)
	b = Y(b, a, c)
	fmt.Println(b)
}

func Y(y, a, c int) int {
	y = a - c
	return y

}
