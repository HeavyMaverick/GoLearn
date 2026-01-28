package another

// // Написать 3 функции
// // writer - генерит числа от 1 до 10
// // doubler - умножает числа на 2, имитируя работу (500мс)
// // reader - читает и выводит на экран

// // pipeline
// func reader(ch <-chan int) {
// 	for v := range ch {
// 		fmt.Println(v)
// 	}
// }
// func doubler(ch <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for v := range ch {
// 			time.Sleep(time.Duration(500) * time.Millisecond)
// 			out <- v * 2
// 		}
// 		close(out)
// 	}()
// 	return out
// }
// func writer() <-chan int {
// 	ch := make(chan int)
// 	go func() {
// 		for i := range 10 {
// 			ch <- i + 1
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }
// func main() {
// 	reader(doubler(writer()))
// }
