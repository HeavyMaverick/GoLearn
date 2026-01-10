package additional

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
