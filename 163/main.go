package main

import (
	"fmt"
	"math"
)

func main() {
	c := 16.0

	f := func(c float64) float64 {
		d := math.Sqrt(c)
		return d
	}(c)

	fmt.Println(f)

}
