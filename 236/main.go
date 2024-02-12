package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	err error
}

func (sq sqrtError) Error() string {
	return fmt.Sprintf("Sqrt error: %v", sq.err)
}

func main() {
	fmt.Println("Input any number")
	fmt.Print(">-")

	var f float64
	_, err := fmt.Scan(&f)
	if err != nil {
		log.Println("You must input a number", err)
		return
	}

	sf, err := sqrt(f)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("The square root of %v is %v\n", f, sf)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		sre := fmt.Errorf("cannot take square root of a negative number %v", f)
		return 0, sqrtError{sre}
	}
	return math.Sqrt(f), nil
}
