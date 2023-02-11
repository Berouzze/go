package main

import "fmt"

func main() {
	var Number int
	fmt.Scan(&Number)
	if Number%7 == 0 && Number%23 == 0 {
		fmt.Print(Number, " Кратно ", "7 и 23")
	} else {
		fmt.Print(Number, " Не кратно ", "7 и 23")
	}
}
