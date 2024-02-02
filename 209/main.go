package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	gr := 100
	wg.Add(gr)
	count := 0

	for i := 0; i < gr; i++ {

		go func() {
			mx.Lock()
			v := count
			v++
			count = v
			fmt.Println(count)
			mx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Count", count)
	fmt.Println("Exiting program")
}
