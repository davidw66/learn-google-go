package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU's", runtime.NumCPU())
	fmt.Println("Go routines", runtime.NumGoroutine())
	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		// fmt.Println("Go routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines", runtime.NumGoroutine())
	fmt.Println("Count:", counter)
}
