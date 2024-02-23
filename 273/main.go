package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := range 10 {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
	fmt.Println("Exiting")
}
