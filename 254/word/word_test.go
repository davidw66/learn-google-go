package word

import (
	"fmt"
	"hello/254/quote"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	ts := "This is exactly six words long"

	got := Count(ts)
	want := 6

	if got != want {
		t.Errorf("\n Got: %v\nWant: %v", got, want)
	}
}

func ExampleCount() {
	ts := "This is exactly six words long"
	fmt.Println(Count(ts))

	// Output:
	// 6
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func TestUseCount(t *testing.T) {
	tq := "one two two three three three"
	got := UseCount(strings.ToLower(tq))

	for k, v := range got {
		switch k {
		case "one":
			if v != 1 {
				t.Errorf("\nFor %v \nGot: %v\nWant: 1", k, v)
			}
		case "two":
			if v != 2 {
				t.Errorf("\nFor %v \nGot: %v\nWant: 2", k, v)
			}
		case "three":
			if v != 3 {
				t.Errorf("\nFor %v \nGot: %v\nWant: 3", k, v)
			}
		}
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
