package another

// var (
// 	n       int
// 	wg      sync.WaitGroup
// 	mu      sync.Mutex
// 	balance int64
// )

// Завершите пример  решения классической гонки данных (data race),
// аналог предыдущей задачи, но уже с использованием  мьютексов.
// Несколько горутин одновременно читают и изменяют переменную balance

// func add() {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	defer wg.Done()
// 	balance += 3
// }

// func main() {
// 	_, err := fmt.Scan(&n)
// 	if err != nil {
// 		log.Println("Error", err)
// 	}
// 	wg.Add(n)
// 	for i := 0; i < n; i++ {
// 		go add()
// 	}
// 	wg.Wait()
// 	fmt.Println("Баланс:", balance)
// }
