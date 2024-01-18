package main

import "fmt"

func main() {
	var a1 = increment()
	a2 := increment()
	a3 := increment()
	fmt.Println("a1 is", a1())
	fmt.Println("a2 is", a2())
	fmt.Println("a3 is", a3())
	fmt.Println("a1 is", a1())
	fmt.Println("a3 is", a3())
	fmt.Println("a1 is", a1())
}

func increment() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
