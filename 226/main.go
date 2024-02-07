package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 100
	}()

	v, ok := <-c
	if ok {
		fmt.Println("channel is open")
		fmt.Println(v, ok)
	}

	close(c)

	v, ok = <-c
	if !ok {
		fmt.Println("channel is closed")
		fmt.Println(v, ok)
	}
}
