package main

import "testing"

func TestAverage(t *testing.T) {
	got := average(2, 3, 4, 1)
	want := 2.5

	if got != want {
		t.Errorf("\n Got: %v\nWant: %v", got, want)
	}
}
