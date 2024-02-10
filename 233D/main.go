package main

import (
	"fmt"
	"log"
	"math"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	s, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math: square root of a negative number: %v", f)
		return 0, norgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return math.Sqrt(f), nil
}
