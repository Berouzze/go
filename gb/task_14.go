package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	c := (b / 100) % 10
	if c == 0 {
		fmt.Print("Nil")
	} else {
		fmt.Print(c)
	}

}
