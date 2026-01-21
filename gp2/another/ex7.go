//	func GenerateSquares(quit <-chan int) <-chan int {
//		squares := make(chan int) // 1 - Создать канал squares
//		go func() {
//			i := 0
//			defer close(squares)
//			for {
//				i += 1
//				select {
//				case squares <- i * i: // 3 - случай отправки квадрата числа в канал
//				case <-quit:
//					return
//				}
//			}
//		}()
//		return squares
//	}
package another
