package main

import "fmt"

// It is convention to use capital T but not necessary
// nor is it necessary to name the function xxxT
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addT(2, 3))
	fmt.Println(addT(4.7, 5.9))
}
