package main

import "fmt"

func main() {
	c := gen(4, 5, 6)
	out := sq(c)

	for numSqr := range out {
		fmt.Println(numSqr)
	}
}

func gen(i ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range i {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
