package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mx sync.Mutex
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")

	wg.Wait()
	fmt.Println("Final Counter: ", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		mx.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, "Counter", counter)
		mx.Unlock()
	}
	wg.Done()
}
