package another

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	// Горутина 1: пытается отправить в ch1, потом прочитать из ch2
// 	go func() {
// 		for {
// 			select {
// 			case ch1 <- 1:
// 				fmt.Println("Горутина 1 отправила в ch1")
// 			case <-ch2:
// 				fmt.Println("Горутина 1 получила из ch2")
// 			}
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()

// 	// Горутина 2: пытается отправить в ch2, потом прочитать из ch1
// 	go func() {
// 		for {
// 			select {
// 			case ch2 <- 2:
// 				fmt.Println("Горутина 2 отправила в ch2")
// 			case <-ch1:
// 				fmt.Println("Горутина 2 получила из ch1")
// 			}
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()

//		// Даем livelock'у поработать
//		time.Sleep(2 * time.Second)
//		fmt.Println("Livelock продолжается...")
//	}
