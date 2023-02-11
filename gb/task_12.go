package main

import (
	"fmt"
)

func main() {
	var s int
	fmt.Scan(&s)

	r := s / 100
	t := s % 10
	fmt.Print(r)
	fmt.Print(t)
}
