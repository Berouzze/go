package main

import "fmt"

func main() {
	var dayNum int
	fmt.Scan(&dayNum)
	if dayNum < 8 && dayNum > 5 {
		fmt.Println("Weekend")
	} else {
		fmt.Println("Workday")
	}
}
