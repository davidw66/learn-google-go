package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go add(c)

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Exiting")
}

func add(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
