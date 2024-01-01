package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	var xn1, xn2, xn3, xn4 []int

	xn1 = append(xn1, xi[:5]...)
	xn2 = append(xn2, xi[5:]...)
	xn3 = append(xn3, xi[2:7]...)
	xn4 = append(xi[1:6], xi[8:10]...)

	fmt.Println(xn1)
	fmt.Println(xn2)
	fmt.Println(xn3)
	fmt.Println(xn4)
}
