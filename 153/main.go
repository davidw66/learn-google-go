package main

import "fmt"

func main() {
	a := foo()
	fmt.Println(a)

	fmt.Println(bar())

}

func foo() int {
	return 23
}

func bar() (int, string) {
	return 23, "and me"
}
