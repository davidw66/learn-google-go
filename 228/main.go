package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	c := make(chan int)

	go addTo(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Exiting")
}

func addTo(c chan<- int) {
	const gRoutine = 10
	var wg sync.WaitGroup
	wg.Add(gRoutine)
	for i := 0; i < gRoutine; i++ {
		go func() {
			for v := 0; v < gRoutine; v++ {
				c <- rand.Intn(1000) + v
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c)
}
