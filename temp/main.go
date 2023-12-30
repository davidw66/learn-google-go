package main

import "fmt"

func main() {
	a := [...]int{10, 12, 14, 16, 18}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
