package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for C := 1; C <= N; C++ {
		if C%2 == 0 {
			fmt.Println(C)
		}
	}
}
