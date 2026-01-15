package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func max(n ...int) (max int) {
	for i := range n {
		if n[i] > max {
			max = n[i]
		}
	}
	return max
}
func unique(s []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, value := range s {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}
func main() {
	var s string
	fmt.Scan(&s)
	for i := range s {

	}
}
