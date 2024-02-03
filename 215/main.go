package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go send(c, xi)

	// did not start goroutine so it would block
	receive(c)

	fmt.Println("Exiting")
}

func send(c chan<- int, xi []int) {
	for _, v := range xi {
		c <- v
	}
	close(c)
}

func receive(c <-chan int) {
	for {
		x, ok := <-c
		if !ok {
			return
		}
		fmt.Println(x)
	}
}
