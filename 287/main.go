package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := incrementor(6)

	totalCount := 0
	for n := range c {
		totalCount++
		fmt.Println(n)
	}

	fmt.Println("Total count", totalCount)
}

func incrementor(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := range n {
		go func(i int) {
			for k := range 20 {
				c <- fmt.Sprint("Process:  "+strconv.Itoa(i)+" printing: ", k)
			}
			done <- true
		}(i)
	}

	go func() {
		for range n {
			<-done
		}
		close(c)
	}()

	return (c)
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
