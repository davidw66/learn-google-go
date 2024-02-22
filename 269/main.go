package main

import "fmt"

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := range 100_000 {
			c <- i
		}
		close(c)
	}()

	for range n {
		go func() {
			for i := range c {
				fmt.Println(i)
			}
			done <- true
		}()
	}

	for range n {
		<-done
	}

	fmt.Println("Exiting")
}
