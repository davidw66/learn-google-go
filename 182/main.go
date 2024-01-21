package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(add(4.7, 5.9))

}

func add[g int | float64](a, b g) g {
	return a + b
}
