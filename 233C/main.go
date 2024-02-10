package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	s, err := sqrt(-42)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: square root of a negative number: %v", f)
	}
	return math.Sqrt(f), nil
}
