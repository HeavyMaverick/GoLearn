package main

import (
	"fmt"
	"strconv"
)

// var strs []string = []string{"flower", "flow", "flight"}
// var strs []string = []string{"dog", "racecar", "car"}
// var strs []string = []string{"cir", "car"}
func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for _, v := range strs {
			if i >= len(v) || v[i] != ch {
				return prefix
			}
		}
		prefix += string(ch)
	}
	return prefix
}

// 38 -> 11 -> 2
func AddDigits(num int) int {
	myStr := strconv.Itoa(num)
	if len(myStr) == 1 {
		return num
	}
	var count int
	for _, ch := range myStr {
		digit, err := strconv.Atoi(string(ch))
		fmt.Println(digit)
		if err != nil {
			fmt.Println(err)
		}
		count += digit
	}
	if len(strconv.Itoa(count)) == 1 {
		return count
	}
	return AddDigits(count)
}
