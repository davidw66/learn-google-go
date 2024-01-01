package main

import "fmt"

func main() {
	// xi := []int{}
	xi := make([]int, 0, 100)
	for c := 42; c <= 51; c++ {
		xi = append(xi, c)
	}

	for i, v := range xi {
		fmt.Printf("Index position %d is of value %d and of type %T\n", i, v, v)
	}
}
