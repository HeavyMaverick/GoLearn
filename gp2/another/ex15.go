package another

import (
	"fmt"
	"sync"
)

// Задача: параллельный подсчёт длины строк
func CountLengths(strs []string) []int {
	wg := &sync.WaitGroup{}
	var result []int = make([]int, len(strs))
	for i, v := range strs {
		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()
			result[i] = len(v)
		}(i, v)
	}
	wg.Wait()
	return result
}

func main1231231232131() {
	slice := []string{"aaa", "bbbb", "ccccc"}
	fmt.Println(CountLengths(slice))
}
