package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			c++
			fmt.Println(c, "\tX is", x)
		}
	}
}
