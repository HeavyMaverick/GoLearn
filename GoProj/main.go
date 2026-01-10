package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
