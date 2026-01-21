package another

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// Код в котором одна горутина генерирует данные,
// а другая горутина выводит эти данные на экран.
// Требуется создать канал ch и горутину для отправки в него строки.
// Основная горутина (main) уже написана (читает из канала и выводит сообщение).
// Sample Input:
// Привет, Go!
// Sample Output:
// Привет, Go!

// func Another() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Scan()
// 	s := sc.Text()
// 	ch := make(chan string)

// 	go func(channel chan<- string) {
// 		channel <- s
// 	}(ch)
// 	msg := <-ch
// 	fmt.Println("Прочитано из канала", msg)
// }
