package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for C := (-N) + 1; C < N; C++ {
		fmt.Println(C)

	}
}
