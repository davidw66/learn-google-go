package main

import "fmt"

func main() {
	avg := average(5, 9, 4, 2)
	fmt.Println(avg)
}

func average(f ...float64) float64 {
	var numerator float64
	var denominator = float64(len(f))

	for _, v := range f {
		numerator += v
	}
	return numerator/denominator + 1
}
