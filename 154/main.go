package main

import "fmt"

func main() {
	xi := []int{5, 6, 9, 8, 12, 45}
	sum := foo(xi...)
	fmt.Println(sum)
	newSum := bar(xi)
	fmt.Println(newSum)

}

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
