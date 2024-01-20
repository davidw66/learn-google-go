package main

import "fmt"

type person struct {
	first string
}

func (p person) changeName(s string) person {
	p.first = s
	return p
}

func main() {
	p1 := person{
		first: "Mary",
	}
	fmt.Println(p1)

	p1 = cn1(p1, "Robert")
	fmt.Println(p1)

	cn2(&p1, "Julia")
	fmt.Println(p1)

	p1 = p1.changeName("Mark")
	fmt.Println(p1)
}

func cn1(p person, s string) person {
	p.first = s
	return p
}

func cn2(p *person, s string) {
	p.first = s
}
