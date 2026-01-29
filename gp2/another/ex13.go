package another

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type outVal struct {
	val int
	err error
}

var errTimeout = errors.New("Error timeout")

func processData(val int, ctx context.Context) chan outVal {
	ch := make(chan struct{})
	out := make(chan outVal)
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(ch)
	}()
	go func() {
		select {
		case <-ch:
			out <- outVal{
				val: val * 2,
				err: nil,
			}
		case <-ctx.Done():
			out <- outVal{
				val: 0,
				err: errTimeout,
			}
		}
	}()
	return out
}

// операция должна выполняться не более 5 секунд
func processParallel(in <-chan int, out chan<- int, numWorkers int, ctx context.Context) {
	wg := &sync.WaitGroup{}
	for range numWorkers {
		wg.Add(1)
		go worker(ctx, in, out, wg)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
}

func main() {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		defer close(in)
		for i := range 100 {
			select {
			case in <- i:
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	now := time.Now()
	processParallel(in, out, 5, ctx)
	for val := range out {
		fmt.Println(val)
	}
	fmt.Println(time.Since(now))
}

func worker(ctx context.Context, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-in:
			if !ok {
				return
			}
			select {
			case ov := <-processData(v, ctx):
				if ov.err != nil {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- ov.val:
				}
			case <-ctx.Done():
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
