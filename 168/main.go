package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timedFunc(work)
}

func timedFunc(fn func()) {
	startJob := time.Now()
	fn()
	endJob := time.Since(startJob)
	fmt.Println("The job took", endJob)
}
func work() {
	time.Sleep(time.Second * time.Duration(rand.Intn(4)))
	fmt.Println("Work completed")
}
