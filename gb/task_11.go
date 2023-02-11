package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int
	a, b := 10, 99
	num = a + rand.Intn(b-a)
	c := num % 10
	d := num / 10
	if c > d {
		fmt.Println(c)
	} else {
		fmt.Println(d)
	}
}
