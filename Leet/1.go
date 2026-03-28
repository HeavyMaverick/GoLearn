package main

import (
	"strconv"
	"strings"
)

func main() {
}

func isPalindrome(x int) bool {
	var builder strings.Builder
	var d string
	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteString(string(s[i]))
	}
	d = builder.String()
	if s == d {
		return true
	} else {
		return false
	}
}
