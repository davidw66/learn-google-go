package main

import "testing"

func TestDoMath(t *testing.T) {
	got := doMath(10, 2, add)
	want := 12

	if got != want {
		t.Errorf("\nFor add:\n got %d want %d\n", got, want)
	}

	got = doMath(10, 2, subtract)
	want = 8

	if got != want {
		t.Errorf("\nFor subtract:\n got %d want %d\n", got, want)
	}

}

func TestAdd(t *testing.T) {
	got := add(5, 2)
	want := 7

	if got != want {
		t.Errorf("\n Got %d\nWant %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(5, 2)
	want := 3

	if got != want {
		t.Errorf("\n Got %d\nWant %d", got, want)
	}
}
