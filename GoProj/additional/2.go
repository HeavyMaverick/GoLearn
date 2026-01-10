// На вход подается число типа float64.
// Необходимо вывести преобразованное число, с 4 цифрами после запятой
package additional

import "fmt"

func Additional() {
	var (
		c float64
	)
	_, err := fmt.Scanf("%f", &c)
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	fmt.Printf("%.4f\n", c)
}
