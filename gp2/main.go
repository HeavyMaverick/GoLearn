// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"
// )

// type JsonData struct {
// 	Rates map[string]float64 `json:"rates"`
// }

// func main() {
// 	client := &http.Client{}
// 	request, err := http.NewRequest("GET", "https://open.er-api.com/v6/latest/USD", nil)
// 	if err != nil {
// 		fmt.Println("Request error", err)
// 	}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println("Response err", err)
// 	}
// 	defer response.Body.Close()

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Reading error", err)
// 	}
// 	fmt.Println(string(body))
// 	var JsonData JsonData
// 	err = json.Unmarshal(body, &JsonData)
// 	if err != nil {
// 		fmt.Println("Unmarshal err", err)
// 	}
// 	fmt.Println(JsonData.Rates)
// 	var counter int = 0
// 	for k := range JsonData.Rates {
// 		if strings.HasPrefix(k, "T") {
// 			counter += 1
// 		}
// 	}
// 	fmt.Println(counter)
// }

// package main

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
package main

import (
	"fmt"
)

// Первая стадия: генерирует числа от 1 до n и отправляет в канал
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Вторая стадия: принимает числа, возводит в квадрат и отправляет дальше
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Третья стадия: принимает числа и выводит их
func print(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	// Создаем конвейер: ген -> sq -> print
	c := gen(2, 3, 4, 5, 9999)
	c2 := sq(c)
	print(c2)
}
