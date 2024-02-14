package main

import "fmt"

func main() {
	fmt.Println("Input your dogs age")
	fmt.Print(">-")

	var a int
	fmt.Scan(&a)
	r := dogYears(a)
	fmt.Printf("Your dog is %v years old", r)
}

func dogYears(i int) int {
	return i * 7
}
