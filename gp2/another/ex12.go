package another

import (
	"fmt"
	"strconv"
)

func countNumbers(nums []int) string {
	var a, b, c int
	for i := range nums {
		if nums[i] > 0 {
			a += 1
		} else {
			if nums[i] < 0 {
				b += 1
			} else {
				if nums[i] == 0 {
					c += 1
				}
			}
		}
	}
	var s string = "выше нуля: " + strconv.Itoa(a) + " ниже нуля: " + strconv.Itoa(b) + ", равна нулю: " + strconv.Itoa(c)
	return s
} // выше нуля: 3, ниже нуля: 2, равна нулю: 2

func main12312() {
	var numsd []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, -1, -2, -3}
	fmt.Println(countNumbers(numsd))
}
