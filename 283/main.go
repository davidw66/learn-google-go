package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	in := gen()
	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	var wg sync.WaitGroup
	const n = 500000
	wg.Add(n)
	out := make(chan int)

	for range n {
		go func() {
			for range 5 {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
				total := 1
				for i := rand.Intn(15) + 3; i > 0; i-- {
					total *= i
				}
				out <- total
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n
		}
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
