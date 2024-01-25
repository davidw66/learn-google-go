package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	xi := []int{12, 4, 723, 96, 1, 14, 16}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Fprintln(os.Stdout, xi)
	sort.Ints(xi)
	fmt.Fprintln(os.Stdout, xi)

	fmt.Println("------------------------")

	fmt.Fprintln(os.Stdout, xs)
	sort.Strings(xs)
	fmt.Fprintln(os.Stdout, xs)

}
