package main

import (
	"fmt"
	"slices"
)

func main() {
	var capacity int
	_, err := fmt.Scan(&capacity)
	if err != nil {
		fmt.Println("Err", err)
	}
	queue := make([]int, 0, capacity)
	for i := 0; i < 10; i++ {
		queue = append(queue, i)
		if i >= len(queue) {
			queue = slices.Delete(queue, 0, 1)
			var newQueue []int
			for _, x := range queue {
				newQueue = append(newQueue, x)
			}
			queue = newQueue
		}
		fmt.Printf("Добавлен %d: %v\n", i, queue)
	}
}
