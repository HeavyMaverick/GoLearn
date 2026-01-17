package additional

import (
	"fmt"
	"strconv"
	"strings"
)

func Fibonacci(n int) []int {
	var slice []int = []int{0, 1}
	if n <= 1 {
		var ans []int = []int{0, 1, 1}
		return ans
	}
	for i := 1; ; i++ {
		if slice[i-1]+slice[i] > n {
			break
		}
		if slice[i-1]+slice[i] <= n {
			slice = append(slice, slice[i-1]+slice[i])
		}
	}
	return slice
}
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
func Max(n ...int) (max int) {
	for i := range n {
		if n[i] > max {
			max = n[i]
		}
	}
	return max
}
func Unique(s []string) []string {
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
func IsPalindrome() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("Error", err)
	}
	var runes []rune = []rune(strings.ToLower(s))
	length := len(runes)
	var ans bool = true
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			ans = false
			break
		}
	}
	fmt.Println(strconv.FormatBool(ans))
}
