package main

import (
	"fmt"
	"math/rand"
)

func main() {
	xi := []int{}
	for range 100 {
		axi := rand.Intn(17) + 3
		xi = append(xi, axi)
	}
	c := proc(xi)
	fact := factorial(c)

	for factorials := range fact {
		fmt.Println("factorial is", factorials)
	}
	fmt.Println("Exiting")
}

func proc(xi []int) chan int {
	out := make(chan int)
	go func() {
		for _, v := range xi {
			out <- v
		}
		close(out)
	}()
	return out
}

func factorial(f chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range f {
			facCalc := 1
			if n <= 0 {
				out <- 0
				continue
			}
			for v := range n + 1 {
				if v == 0 {
					continue
				}
				facCalc *= v
			}
			out <- facCalc
		}
		close(out)
	}()
	return out
}
