package patterns

import (
	"context"
	"sync"
)

// Generate
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// FanIN pattern
func fanIn(ctx context.Context, chans []chan int) chan int {
	out := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				select {
				case v, ok := <-ch:
					if !ok {
						return
					}
					select {
					case out <- v:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}()
		}
		close(out)
		wg.Wait()
	}()
	return out
}
