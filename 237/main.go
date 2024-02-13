package main

import (
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("Math error at lat %v, long %v: error %v", se.lat, se.long, se.err)
}

func main() {
	fmt.Println("Input any number")
	fmt.Print(">-")

	var f float64
	_, err := fmt.Scan(&f)
	if err != nil {
		fmt.Println(err)
		return
	}

	sq, err := sqrt(f)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("The square root of %v is %v\n", f, sq)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		se := fmt.Errorf("cannot take square root of negative number %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", se}
	}
	return math.Sqrt(f), nil
}
