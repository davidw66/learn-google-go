package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		input []int
		want  float64
	}

	tests := []test{
		{[]int{-1, 2, 3, 4, 5}, 3},
		{[]int{3, 6, 9, 8}, 7},
		{[]int{100000, 1, 3, 4, 5}, 4},
		{[]int{3, 4, 4, 5}, 4},
		{[]int{9, 2, 4, 1}, 3},
		{[]int{9, 2, 7, 1}, 4.5},
	}

	for _, v := range tests {
		got := CenteredAvg(v.input)

		if got != v.want {
			t.Errorf("\n Got: %v\nWant: %v\n", got, v.want)
		}
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5}

	fmt.Println(CenteredAvg(xi))
	// Output:
	// 3
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{5, 9, 5, 4, 8, 9, 3, 6, 55, 2, 1})
	}
}
