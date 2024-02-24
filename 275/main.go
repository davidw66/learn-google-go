package main

import "fmt"

func main() {
	for numSqr := range sq(gen(4, 5, 6)) {
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
