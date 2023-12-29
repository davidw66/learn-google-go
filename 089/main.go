package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
		"Robert":     89,
		"Mark":       23,
	}

	for k, v := range m {
		fmt.Printf("%s is %d\n", k, v)
	}

	m1 := m["James"]
	fmt.Println(m1)

	m2 := m["Q"]
	fmt.Println(m2)

	v, ok := m["Q"]
	if ok {
		fmt.Println("Q is", v)
	} else {
		fmt.Println("Q is not in the database")
	}

	v, ok = m["James"]
	if ok {
		fmt.Println("James is", v)
	} else {
		fmt.Println("James is not in the database")
	}
	fmt.Println(v, ok)

	if v, ok = m["Robert"]; ok {
		fmt.Println("Robert is", v)
	} else {
		fmt.Println("Robert is not in database")
	}

	if v, ok = m["James"]; ok {
		fmt.Println("James is", v)
	} else {
		fmt.Println("James is not in database")
	}

	a := "Mark"

	if v, ok = m[a]; ok {
		fmt.Println(a, "is", v)
	} else {
		fmt.Println(a, "is not in the database")
	}

}
