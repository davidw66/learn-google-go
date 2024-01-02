package main

import "fmt"

func main() {
	xj := []string{"James", "Bond", "Shaken not stirred"}
	xm := []string{"Miss", "Moneypenny", "I'm 008"}
	xss := [][]string{xj, xm}

	fmt.Println(xss)

	for _, v := range xss {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
