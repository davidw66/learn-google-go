// Package sum adds integers
package sum

// SumInt accepts an unlimited number if of integers and returns the sum of those integers
func SumInt(i ...int) int {
	sumTotal := 0

	for _, v := range i {
		sumTotal += v
	}
	return sumTotal
}
