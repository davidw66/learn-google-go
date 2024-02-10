package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrNorgateMath = errors.New("norgate math: square root of a negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	s, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}
