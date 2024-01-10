package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(loopFactorial(4))

}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)
}

func loopFactorial(i int) int {
	n := 1
	for i > 0 {
		n *= i
		i--
	}
	return n
}
