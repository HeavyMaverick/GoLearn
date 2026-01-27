package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Задача: необходимо уложиться в общее выполнение программы
//5 секунд на всю программу
//Каждая горутина должна спать от 0 до 5 секунд (рандомно).
//Вывести сколько всего времени спали горутины
//и сколько времени заняло выполнение программы.

var maxWaitSeconds = 5

func randomWait() int {
	workSeconds := rand.Intn(5 + 1)
	time.Sleep(time.Duration(workSeconds) * time.Second)
	return workSeconds
}
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
