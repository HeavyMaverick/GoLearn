package main

func main() {
	mainSeconds := 0
	totalWorkSecond := 0
	timeNow := time.Now()
	ch := make(chan int, 100)
	for range 100 {
		go func() {
			ch <- randomWait()
		}()
	}
	for range 100 {
		totalWorkSecond += <-ch
	}
	mainSeconds = int(time.Since(timeNow).Seconds())
	fmt.Println("main", mainSeconds)
	fmt.Println("total", totalWorkSecond)
}
