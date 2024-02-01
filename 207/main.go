package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go iCount()
	go betWhere()

	wg.Wait()
	fmt.Println("Exiting")
}

func iCount() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Microsecond * 1)
	}
	wg.Done()
}

func betWhere() {
	for i := 0; i < 4; i++ {
		fmt.Println("Where will I show up")
		time.Sleep(time.Microsecond * 3)
	}
	wg.Done()
}
