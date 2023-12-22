package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
