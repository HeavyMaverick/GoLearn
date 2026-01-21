package another

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// 	"sync/atomic"
// )

// Завершите пример  решения классической гонки данных (data race)
// с использованием  атомарных операций и WaitGroup.
// Несколько горутин одновременно читают и изменяют переменную balance.
// Напишите код функции add которая увеличивает общую переменную balance на 3.
// Ключевые моменты:
// atomic работает только с указателями.
// Нужна синхронизация (WaitGroup) для ожидания горутин.
// Передавать указатель на общую переменную.
// Sample Input:
// 10
// Sample Output:
// Баланс: 30
// func add(bal *int64, wg *sync.WaitGroup) {
// 	atomic.AddInt64(bal, 3)
// 	wg.Done()
// }

// func main() {
// 	var n int
// 	_, err := fmt.Scan(&n)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	var balance int64 = 0
// 	var wg sync.WaitGroup

// 	wg.Add(n)

// 	for i := 0; i < n; i++ {
// 		go add(&balance, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("Баланс:", balance)
// }
