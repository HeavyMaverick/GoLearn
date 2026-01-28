package another

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// имеется функция, которая работает неопределенно долго (до 100 секунд)
func randomTimeWork(done chan struct{}) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
	close(done)
}

// Написать обертку для этой функции, которая будет прерывать выполнение, если
// функция работает дольше 3х секунд, и возвращать ошибку
func predictableTimeWork(ctx context.Context) {
	done := make(chan struct{})
	go randomTimeWork(done)
	select {
	case <-done:
		fmt.Println("Выход по завершению работы")
	case <-ctx.Done():
		fmt.Println("Выход по контексту")
		return
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	predictableTimeWork(ctx)
}
