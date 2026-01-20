package another

// import (
// 	"fmt"
// )

// // Первая стадия: генерирует числа от 1 до n и отправляет в канал
// func gen(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// // Вторая стадия: принимает числа, возводит в квадрат и отправляет дальше
// func sq(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// // Третья стадия: принимает числа и выводит их
// func print(in <-chan int) {
// 	for n := range in {
// 		fmt.Println(n)
// 	}
// }

// func main() {
// 	// Создаем конвейер: ген -> sq -> print
// 	c := gen(2, 3, 4, 5, 9999)
// 	c2 := sq(c)
// 	print(c2)
// }
