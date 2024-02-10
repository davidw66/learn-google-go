package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	s, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of a negative number")
	}
	return math.Sqrt(f), nil
}
