package main

import "fmt"

// This is called a type set
type myNumbers interface {
	int | float64
}

func AddT[T myNumbers](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	i := AddT(4, 5, 6)
	fmt.Println(i)

	f := AddT(6.9, 7.6, 1.2)
	fmt.Println(f)
}
