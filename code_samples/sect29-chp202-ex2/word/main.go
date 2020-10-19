// Package word provides custom functions that returns info on strings.
package word

import "strings"

// UseCount returns the count of each word in the string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in the string.
func Count(s string) int {
	xs := strings.Fields(s)

	return len(xs)
}
