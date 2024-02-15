package main

import "testing"

func TestAverage(t *testing.T) {
	type test struct {
		data []float64
		want float64
	}
	tests := []test{
		{[]float64{2, 2}, 2},
		{[]float64{2, 2, 8}, 4},
		{[]float64{2, -2, 15}, 5},
		{[]float64{6, 3, 8, -7}, 2.5},
		{[]float64{3, 4, 8}, 5},
	}

	for _, v := range tests {
		got := average(v.data...)
		want := v.want

		if got != want {
			t.Errorf("\n Got: %v\nWant: %v\n", got, want)
		}
	}
}
