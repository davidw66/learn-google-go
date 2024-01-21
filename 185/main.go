package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// This is called a type set
type myNumbers interface {
	constraints.Integer | constraints.Float
}

func AddT[T myNumbers](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

type myAlias int

func main() {
	var n myAlias = 42

	i := AddT(4, 5, 6, n)
	fmt.Println(i)

	f := AddT(6.9, 7.6, 1.2)
	fmt.Println(f)
}
