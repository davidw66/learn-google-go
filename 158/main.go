package main

import "fmt"

func main() {
	tot := add(5, 8)
	fmt.Println(tot)
}

func add(a int, b int) int {
	return a + b
}
