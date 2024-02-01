package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type greeting interface {
	hello()
	goodbye()
}

func (p *person) hello() {
	fmt.Println("Hello", p.first, p.last)
}

func (p person) goodbye() {
	fmt.Println("Goodbye", p.first, p.last)
}

func sayHello(g greeting) {
	g.hello()
}

func sayGoodbye(g greeting) {
	g.goodbye()
}

func main() {
	matt := person{
		first: "Matt",
		last:  "Zinger",
	}

	matt.hello()
	matt.goodbye()

	fmt.Println("---------------------------")

	sayHello(&matt)
	sayGoodbye(&matt)
}
