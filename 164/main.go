package main

import "fmt"

func main() {
	x := func() {
		for i := 0; i < 5; i++ {
			fmt.Println("I is", i)
		}
	}

	x()
	fmt.Println("****************")
	y()
}

var y = func() {
	for c := 1; c < 5; c++ {
		fmt.Println("C is", c)
	}
}
