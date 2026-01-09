// На вход подается число типа float64.
// Необходимо вывести преобразованное число, с 4 цифрами после запятой
// 2 знака после запятой (fmt.Printf("%.2f\n", num)).
package main

import "fmt"

var (
	c float64
)

func main() {
	_, err := fmt.Scanf("%f", &c)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Printf("%.4f\n", c)
}
