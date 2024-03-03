package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// It is idiomatic to name errors errName or ErrName
var ErrNorgateMath = errors.New("norgate math: cannot take square root of negative number")

func main() {
	var n float64
	fmt.Println("Input any positive number")
	fmt.Print(">-\t")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Please input a valid number")
	}
	sqrtN, err := sqrt(n)
	if err != nil {
		log.Fatal(err, n)
	}
	fmt.Println("The square root of", n, "is", sqrtN)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}
