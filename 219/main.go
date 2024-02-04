package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	fanin := make(chan int)

	go send(even, odd, quit)

	go receive(even, odd, quit, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Exiting")
}

func send(even, odd, quit chan<- int) {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
	close(quit)
}

/*
//Instructors way - remove quit chan to use
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
*/

func receive(even, odd, quit <-chan int, fanin chan<- int) {
	for {
		select {
		case v := <-even:
			fanin <- v
		case v := <-odd:
			fanin <- v
		case _, ok := <-quit:
			if !ok {
				close(fanin)
				return
			}
		}
	}
}
