package main

import "fmt"

type Number interface {
	~int64 | ~float64
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  34.99,
		"second": 12.97,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
		SumNumbers[string, int64](ints),
		SumNumbers[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
