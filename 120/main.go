package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":       {"shaken not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"cats", "ice cream", "sunsets"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	// m["bond_james"] = jFav
	// fmt.Println(m["bond_james"])
	fmt.Println("-----------------")

	for k, v := range m {
		fmt.Println(k, "favorite things are:")
		for i, c := range v {
			fmt.Println(i, c)
		}
		fmt.Println("-----------------")
	}

	if _, ok := m["moneypenny_jenny"]; !ok {
		fmt.Println("No record exists to delete")
	} else {
		delete(m, "moneypenny_jenny")
	}

	fmt.Println("*****************")

	for k, v := range m {
		fmt.Println(k, "favorite things are:")
		for i, c := range v {
			fmt.Println(i, c)
		}
		fmt.Println("*****************")
	}
	/*
	   mb := make(map[string][]string)
	   mb["bond_james"] = []string{"shaken not stirred", "martinis", "fast cars"}
	   mb["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	   mb["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	   mb["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	   	for k, v := range mb {
	   		fmt.Println(k)
	   		for i, c := range v {
	   			fmt.Println(i, c)
	   		}
	   	}
	*/
}
