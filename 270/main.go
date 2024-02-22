package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := range 10 {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	factorial := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		factorial <- sum
		close(factorial)
	}()
	return factorial
}
