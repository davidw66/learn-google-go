package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()

	fmt.Println("Exiting program")
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}

	wg.Done()
}
