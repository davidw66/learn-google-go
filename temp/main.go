package main

import "fmt"

func main() {
	a := []int{10, 12, 14, 16, 18}
	b := a
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("------------------")
	// notice that a and b both change
	a[0] = 4
	fmt.Println("a", a)
	fmt.Println("b", b)

	fmt.Println("------------------")
	c := []int{10, 12, 14, 16, 18}
	d := make([]int, 5, 100)
	copy(d, c)
	fmt.Println("c", c)
	fmt.Println("d", d)

	fmt.Println("------------------")
	// notice that only slice c changes
	c[0] = 4
	fmt.Println("c", c)
	fmt.Println("d", d)
}
