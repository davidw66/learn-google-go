package main

import "fmt"

func main() {
	c := incrementor()
	cSum := puller(c)

	fmt.Println(<-cSum)
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := range 10 {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	sumC := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		sumC <- sum
		close(sumC)
	}()
	return sumC
}
