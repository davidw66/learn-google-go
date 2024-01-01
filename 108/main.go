package main

import "fmt"

func main() {
	a := [5]int{}
	for c := 1; c <= 5; c++ {
		a[c-1] = c
	}
	for _, v := range a {
		fmt.Printf("%d is of type %T\n", v, v)
	}

	fmt.Printf("%v is of type %T\n", a, a)
}
