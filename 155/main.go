package main

import "fmt"

func main() {
	defer printMe()
	defer fmt.Println("The second deferred")
	defer fmt.Println("the third deferred")
	fmt.Println("Last line of func main")
}

func printMe() {
	fmt.Println("The first deferred")
}
