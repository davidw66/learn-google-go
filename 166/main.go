package main

import (
	"fmt"
	"math/rand"
)

func main() {

	dogYears := func(i int) int {
		d := i * 7
		return d
	}
	myName := []string{"Fido", "Max", "Boomer", "Molly", "Chewy", "Genghis"}

	name(myName[rand.Intn(6)], dogYears, rand.Intn(10)+1)
}

// Also works
// func dogYears(i int) int {
// 	d := i * 7
// 	return d
// }

func name(s string, f func(int) int, i int) {
	fmt.Printf("My name is %s and I am %d years old.\n", s, f(i))
}
