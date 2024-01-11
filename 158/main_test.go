package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(3, 2)
	want := 5

	if got != want {
		t.Errorf("\n got %d\n want %d", got, want)
	}
}
