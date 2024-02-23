package main

import "fmt"

func main() {
	f := factorial(4)

	for n := range f {
		fmt.Println("Total", n)
	}
}

func factorial(n int) chan int {
	c := make(chan int)
	total := 1

	for i := n; i > 0; i-- {
		total *= i
	}

	go func() {
		c <- total
		close(c)
	}()
	return c
}
