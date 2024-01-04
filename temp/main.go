package main

import "fmt"

type dogYears int

func main() {

	var age dogYears = 4
	fluffy := 4
	age.dogAge(fluffy)
	{
		type person struct {
			first string
			last  string
			age   int
		}

		type secretAgent struct {
			person
			ltk   bool
			first string
		}

		sa1 := secretAgent{
			person: person{
				first: "James",
				last:  "Bond",
				age:   32,
			},
			ltk:   true,
			first: "Ethan",
		}

		p1 := person{
			first: "Jenny",
			last:  "Moneypenny",
			age:   22,
		}

		p3 := struct {
			first string
			last  string
			age   int
		}{
			first: "David",
			last:  "Williamson",
			age:   31,
		}

		fmt.Println(sa1)
		fmt.Println(p1)
		fmt.Println(p3)
		fmt.Println(sa1.first, sa1.first, sa1.person.first)
	}
}

func (d dogYears) dogAge(i int) {
	fmt.Println(int(d) * i)
}
