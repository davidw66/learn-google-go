package main

import (
	"fmt"
	"hello/244/dog"
)

type myDog struct {
	name string
	age  int
}

func main() {
	fmt.Println("Input your dogs name")
	fmt.Print(">-")

	// receive input of dogs name
	var dogName string
	fmt.Scan(&dogName)

	fmt.Println("Input your dogs age in human years")
	fmt.Print(">-")

	// Get input of dogs age in human years
	var dogAge int
	fmt.Scan(&dogAge)

	d1 := myDog{
		name: dogName,
		age:  dog.DogYears(dogAge),
	}

	fmt.Printf("%s's biological age is %d dog years.\n", d1.name, d1.age)
}
