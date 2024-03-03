package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	n := []float64{7, -6.0, 12, -29, 569, 762, -9}
	for _, v := range n {
		s, err := sqrt(v)
		if err != nil {
			log.Println("Math error: ", err)
			continue
		}
		fmt.Printf("The square root of %f is %f\n", v, s)
	}
}

func sqrt(f float64) (float64, error) {
	errNorgateMath := fmt.Errorf("cannot take square root of %f because it is a negative number", f)
	if f < 0 {
		return 0, errNorgateMath
	}
	sqrtN := math.Sqrt(f)
	return sqrtN, nil
}
