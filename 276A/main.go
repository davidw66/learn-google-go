package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := proc()
	fact := factorial(c)

	for factorials := range fact {
		fmt.Println(factorials)
	}
	fmt.Println("Exiting")
}

func proc() chan int {
	out := make(chan int)
	go func() {
		for range 100 {
			out <- rand.Intn(10) + 3
		}
		close(out)
	}()
	return out
}

func factorial(f chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range f {
			out <- facCalc(n)
		}
		close(out)
	}()
	return out
}

func facCalc(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
