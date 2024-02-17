package main

import "fmt"

func main() {
	sumTotal := SumFloat(1, 2, 3)

	fmt.Println(sumTotal)
}

// This does some stuff
func SumFloat(f ...float64) float64 {
	sum_Float := 0.00
	for _, v := range f {
		sum_Float += v
	}
	return sum_Float
}
