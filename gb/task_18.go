package main

import "fmt"

func main() {
	var max, x, min, AA, num, p int
	var y []int
	fmt.Scan(&p)
	for c := 0; c < p; c++ {
		fmt.Scan(&x)
		y = append(y, x)
		AA += x
		if x > max {
			max = x
		}
	}
	min = y[0]
	for idx := range y {
		if y[idx] < min {
			min = y[idx]
		}
	}
	num = max - min
	AA = AA / p
	fmt.Println("Среднее арифметическое ", AA, ", Разница максимально и минимального ", num)
}
