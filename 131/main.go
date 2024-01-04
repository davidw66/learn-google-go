package main

import "fmt"

func main() {
	myFriend := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Robert",
		friends: map[string]int{
			"John":  12,
			"David": 10,
		},
		favDrinks: []string{"Root Beer", "Ginger Beer"},
	}

	fmt.Println(myFriend.first, myFriend.friends["John"], myFriend.favDrinks[1])
}
