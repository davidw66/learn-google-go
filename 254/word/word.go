// Package word contains custom functions for working with strings.
package word

import "strings"

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder

// UseCount analyzes the use of words in strings and returns the word count for each word used.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count accepts the parameter of a string and returns the number of words in the string
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)

}
