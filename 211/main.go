package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	gr := 100
	wg.Add(gr)

	var count atomic.Int64
	// var counter int64

	for i := 0; i < gr; i++ {
		go func() {
			// Will also work
			// atomic.AddInt64(&counter, 1)
			// fmt.Println(atomic.LoadInt64(&counter))

			//Standard library recommended using this instead of atomic.AddInt64(&counter, 1)
			count.Add(1)
			fmt.Println(count.Load())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Count", count.Load())
	// fmt.Println("Count", counter)
	fmt.Println("Exiting program")
}
