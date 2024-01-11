package main

import "fmt"

type dog struct {
	name string
	age  int
}

func (d dog) dogAge() string {

	myDog := fmt.Sprint(d.name, " is ", d.age*7, " dog years old")
	return myDog
}

func main() {
	fido := dog{
		name: "Fido",
		age:  4,
	}

	fmt.Println(fido.dogAge())
}
