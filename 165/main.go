package main

import "fmt"

func main() {
	c := incrementor()
	fmt.Println("This is c", c())
	c()
	fmt.Println("This is c", c())

	d := incrementor()
	d()
	fmt.Println("This is d", d())
	fmt.Println("This is c", c())
}

func incrementor() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}

}
