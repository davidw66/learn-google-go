package main

import "testing"

func TestDogAge(t *testing.T) {
	spot := dog{
		name: "Spot",
		age:  7,
	}
	got := spot.dogAge()
	want := "Spot is 49 dog years old"

	if got != want {
		t.Errorf("\n Got %s\nWant %s\n", got, want)
	}

}
