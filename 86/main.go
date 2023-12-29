package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Series", i)
		for c := 0; c < 5; c++ {
			fmt.Println("\tRound", c)
		}
	}
}
