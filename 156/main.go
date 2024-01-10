package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %s and I am %d years old.\n", p.first, p.age)
}

func main() {
	david := person{
		first: "David",
		age:   25,
	}
	david.speak()
}
