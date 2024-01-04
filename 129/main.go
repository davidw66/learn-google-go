package main

import (
	"fmt"
)

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

func main() {

	p1 := person{
		firstName:  "David",
		lastName:   "Williamson",
		favFlavors: []string{"chocolate", "peanut butter", "mint chocolate chip"},
	}

	marieFlavors := []string{"vanilla", "mango", "water mellon", "chunky monkey"}

	p2 := person{
		firstName:  "Marie",
		lastName:   "Velasco",
		favFlavors: marieFlavors,
	}

	people := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range people {
		v.favorites()
	}

}
func (p person) favorites() {
	fmt.Printf("%s %s's favorite flavors are:\n", p.firstName, p.lastName)
	for i, v := range p.favFlavors {
		fmt.Println(i+1, v)
	}
}
