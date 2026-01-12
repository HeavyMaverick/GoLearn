package additional

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// // Напишите программу, принимающую на вход целое положительное число n
// // и последовательность из n целых положительных чисел
// // и которая выводит в консоль слайс, содержащий только четные числа.
// var (
// 	scanner = bufio.NewScanner(os.Stdin)
// )

// func main() {
// 	scanner.Split(bufio.ScanWords)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())
// 	slice := make([]int, 0)
// 	for range n {
// 		scanner.Scan()
// 		item, _ := strconv.Atoi(scanner.Text())
// 		slice = append(slice, item)
// 	}
// 	var resultSlice = []int{}
// 	for i := 0; i < len(slice); i++ {
// 		if slice[i]%2 == 0 {
// 			resultSlice = append(resultSlice, slice[i])
// 		}
// 	}
// 	fmt.Println(resultSlice)
// }

// На стандартный ввод подаются три целых положительных числа a, b, n, где a < b < n ,
// и n целых положительных чисел, которые должны быть записаны в слайс
// (т.е. необходимо сначала заполнить слайс переданными числами,
// а потом выполнять по порядку каждый пункт с выводом результата).
// Выполните следующие действия со слайсом:
// 1) Присвойте элементу с индексом a  значение 1000 и выведите все элементы слайса.
// 2) Добавьте в конец слайса элементы a и b и выведите все элементы слайса.
// 3) Выведите сумму всех элементов слайса.
// 4) Выведите последние a элементов.
// 5) Выведите все элементы, начиная с элемента с  индексом a и заканчивая элементом с индексом b - 1.
// var (
// 	a, b, n int
// 	numbers = make([]int, 0)
// )

// func main() {
// 	_, err := fmt.Scanf("%d %d %d", &a, &b, &n)
// 	if err != nil {
// 		fmt.Println("Ошибка чтения a, b, n:", err)
// 		return
// 	}
// 	for i := 0; i < n; i++ {
// 		var num int
// 		_, err := fmt.Scan(&num)
// 		if err != nil {
// 			fmt.Printf("Ошибка чтения числа %d: %v\n", i+1, err)
// 			return
// 		}
// 		numbers = append(numbers, num)
// 	}
// 	numbers[a] = 1000
// 	fmt.Println(numbers)
// 	numbers = append(numbers, a, b)
// 	fmt.Println(numbers)
// 	var sum int
// 	for i := 0; i < len(numbers); i++ {
// 		sum += numbers[i]
// 	}
// 	fmt.Println(sum)
// 	fmt.Println(numbers[len(numbers)-a:])
// 	fmt.Println(numbers[a:b])
// }

// Напишите программу,
// // принимающую на вход целое положительное число n
// // и две последовательности изn целых чисел,
// // которая выводит строку,
// // содержащую сумму элементов с одинаковыми индексами
// // (смотрите пример выходных данных).
// // 3
// // 1 2 3
// // 1 2 3
// // output
// // 2 4 6
// import (
// 	"fmt"
// )

// var (
// 	n           int
// 	slice1      = []int{}
// 	slice2      = []int{}
// 	resultSlice = []int{}
// )

// func main() {
// 	_, err := fmt.Scanf("%d", &n)
// 	if err != nil {
// 		fmt.Println("error", err)
// 		return
// 	}
// 	for i := 0; i < n; i++ {
// 		var num int
// 		_, err := fmt.Scan(&num)
// 		if err != nil {
// 			fmt.Println("error", err)
// 			return
// 		}
// 		slice1 = append(slice1, num)
// 	}
// 	for i := 0; i < n; i++ {
// 		var num int
// 		_, err := fmt.Scan(&num)
// 		if err != nil {
// 			fmt.Println("error", err)
// 			return
// 		}
// 		slice2 = append(slice2, num)
// 	}
// 	for i := 0; i < n; i++ {
// 		num := slice1[i] + slice2[i]
// 		resultSlice = append(resultSlice, num)
// 	}
// 	for _, v := range resultSlice {
// 		fmt.Print(v, " ")
// 	}
// }

// // На вход подаётся целое число n и последовательность из n целых чисел.
// // Требуется вывести эту последовательность в обратном порядке.
// func main() {
// 	var n, num int
// 	var slice = []int{}
// 	_, err := fmt.Scan(&n)
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	for i := 0; i < n; i++ {
// 		_, err := fmt.Scan(&num)
// 		if err != nil {
// 			fmt.Println("Error", err)
// 		}
// 		slice = append(slice, num)
// 	}
// 	slice2 := []int{}
// 	for i := len(slice); i > 0; i-- {
// 		slice2 = append(slice2, slice[i-1])
// 	}
// 	for _, v := range slice2 {
// 		fmt.Print(v, " ")
// 	}
// }

// Напишите программу,
// принимающую на вход целое положительное число n
// и последовательность изn целых положительных чисел,
// которая выводит из этой последовательности только уникальные элементы
// (уникальные элементы должны сохранять свой исходный порядок в последовательности).
// func main() {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	scanner.Split(bufio.ScanWords)
// 	scanner.Scan()
// 	n, err := strconv.Atoi(scanner.Text())
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	myMap := map[int]int{}
// 	for range n {
// 		scanner.Scan()
// 		item, _ := strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			fmt.Println("Error in range", err)
// 		}
// 		myMap[item] = n
// 	}
// 	arr := []int{}
// 	for key := range myMap {
// 		arr = append(arr, key)
// 	}
// 	for i := range arr {
// 		fmt.Print(arr[i], " ")
// 	}
// }

// Напишите программу, которая будет запрашивать у пользователя
// целое число n и последовательность из n целых положительных чисел
// и выводить на экран наименьшее и наибольшее число в этой последовательности.
// func main() {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	scanner.Split(bufio.ScanWords)
// 	scanner.Scan()
// 	n, _ := strconv.Atoi(scanner.Text())
// 	var min = math.MaxInt
// 	var max = 0
// 	for range n {
// 		scanner.Scan()
// 		num, _ := strconv.Atoi(scanner.Text())
// 		if num > max {
// 			max = num
// 		}
// 		if num < min {
// 			min = num
// 		}
// 	}
// 	fmt.Println("Наименьшее число:", min)
// 	fmt.Println("Наибольшее число:", max)
// }
