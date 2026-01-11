package additional

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

// Чтение одной строки
func Additional2() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите ваше имя: ")
	if scanner.Scan() {
		name := scanner.Text()
		fmt.Printf("Привет %s!", name)
	}
}

func Additional3() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Вводите несколько строк или введите exit")
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToLower(input) == "exit" {
			break
		}
		fmt.Println("Вы ввели: ", input)
	}
}

func Additional4() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	fmt.Println("Введите несколько слов: ")

	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println("Слово: ", word)
	}
}

/*
Напишите программу,
которая принимает на вход возраст человека и выводит сообщение "Вы совершеннолетний",
если возраст больше или равен 18,
и "Вы несовершеннолетний", если возраст меньше 18.
*/

func Additional5() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	s, err := strconv.Atoi(input)
	if err == nil {
		if s >= 18 {
			fmt.Println("Вы совершеннолетний")
		} else {
			fmt.Println("Вы несовершенолетний")
		}
	} else {
		fmt.Println("Ошибка ", err)
	}
}

/*
Напишите программу, которая принимает на вход два числа и выводит сообщение
"Первое число больше второго", если первое число больше второго,
"Второе число больше первого", если второе число больше первого, и
"Числа равны ", если числа равны.
Числа могут быть с плавающей точкой
*/

func Additional6() {
	var (
		a, b float64
	)
	// n, err := fmt.Scanf("%f %f", &a, &b)
	// if err != nil || n != 2 {
	// 	fmt.Println("Ошибка ", err)
	// }
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	a, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	b, _ = strconv.ParseFloat(scanner.Text(), 64)
	if a > b {
		fmt.Println("Первое число больше второго")
	} else if a < b {
		fmt.Println("Второе число больше первого")
	} else { // a == b
		fmt.Println("Числа равны")
	}
}

/*
Напишите программу,
которая принимает на вход число типа float и проверяет,
является ли оно положительным,
отрицательным или нулем.
Выведите соответствующее сообщение.
*/

func Additional7() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		var a, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Ошибка", err)
		}
		if a > 0 {
			fmt.Println("Положительное число")
		} else if a < 0 {
			fmt.Println("Отрицательное число")
		} else {
			fmt.Println("Нулевое число")
		}
	}
}

func Additional8() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	a64, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	b64, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	op := scanner.Text()
	a := float32(a64)
	b := float32(b64)
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("Деление на ноль")
		}
	default:
		fmt.Println("Некорректное значение!")
	}
}

// Модифицируйте предыдущий вывод "ёлочкой" так,
// чтобы в каждой нечетной строке выводились только нечетные числа, а в каждой четной только четные.
func Additional15() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите число n: ")
	scanner.Scan()
	input := scanner.Text()
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			for j := 1; j <= i; j += 2 {
				fmt.Print(j, " ")
			}
		} else {
			for j := 2; j <= i; j += 2 {
				fmt.Print(j, " ")
			}
		}
		fmt.Println()
	}
}

// На вход подается целое положительное число n, выведете на экран последовательность от 1 до n "ёлочкой".
func Additional14() {
	scanner.Scan()
	input := scanner.Text()
	n64, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println("error -> ", err)
	}
	for i := 1; i <= int(n64); i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

// Напишите программу, которая в бесконечном цикле запрашивает у пользователя число n
// и если на вход пришел 0, программа выводит сумму введенных чисел и завершается.
// Используйте оператор break для выхода из цикла, 0 - условие выхода из бесконечного цикла.
func Additional13() {
	fmt.Println("Введите числа, 0 для выхода")
	var n int64
	for {
		scanner.Scan()
		input := scanner.Text()
		if input == "0" {
			break
		}
		number, err := strconv.ParseInt(input, 10, 64)
		if err == nil {
			n += number
		} else {
			fmt.Println("err -> ", err)
			break
		}
	}
	fmt.Println(n)
}

// Напишите программу, которая принимает на вход число n  и выводит все числа от 1 до n включительно, которые не делятся на 3.
// Используйте оператор continue для пропуска чисел, которые делятся на 3
func Additional12() {
	scanner.Scan()
	input := scanner.Text()
	n64, _ := strconv.ParseInt(input, 10, 64)
	n := int(n64)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// Напишите программу, которая принимает на вход число n и выводит на экран все четные числа от 1 до n включительно.
func Additional11() {
	scanner.Scan()
	input := scanner.Text()
	n64, _ := strconv.ParseInt(input, 10, 64)
	n := int(n64)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// Напишите программу, которая принимает на вход число n и выводит числа от 0 до n не включительно.
func Additional9() {
	scanner.Scan()
	n64, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	n := int(n64)
	for i := range n {
		fmt.Println(i)
	}
}

// Напишите программу, которая принимает на вход число n и считает сумму всех чисел от 1 до n включительно.
func Additional10() {
	scanner.Scan()
	input := scanner.Text()
	n64, _ := strconv.ParseInt(input, 10, 64)
	n := int(n64)
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	fmt.Println(result)
}
