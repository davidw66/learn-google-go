package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	want := 49
	got := dogYears(7)

	if got != want {
		t.Errorf("\n Got %v\nWant %v \n", got, want)
	}
}
