package main

import "fmt"

func main() {
	xi := []int{12, 14, 16, 18, 22, 24}
	c := make(chan int)

	go func() {
		for _, v := range xi {
			c <- v
		}
		close(c)
	}()

	for {
		x, ok := <-c
		if !ok {
			return
		}
		fmt.Println(x)
	}

}
