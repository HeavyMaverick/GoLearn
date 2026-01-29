package another

import (
	"fmt"
	"sync"
)

// Задача: параллельное возведение в квадрат
type chNum struct {
	index int
	value int
}

func squareAll(nums []int) []int {
	result := make([]int, len(nums))
	in := make(chan chNum)
	wg := &sync.WaitGroup{}
	// producer
	go func() {
		defer close(in)
		for index, value := range nums {
			in <- chNum{
				index: index,
				value: value,
			}
		}
	}()
	for j := range in {
		wg.Add(1)
		go func(j chNum) {
			defer wg.Done()
			result[j.index] = j.value * j.value
		}(j)
	}
	wg.Wait()
	return result
}

func main13213123() {
	var nums []int
	for i := range 100 {
		nums = append(nums, i+1)
	}
	result := squareAll(nums)
	fmt.Println(result)
}
