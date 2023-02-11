package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")

	}
}
