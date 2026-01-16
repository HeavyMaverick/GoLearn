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

// Напишите функцию factorial которая принимает неопределенное количество аргументов
// (целые положительные числа) и возвращает факториал суммы переданных значений.
// Т.е. надо вычислить сумму аргументов и вернуть факториал этой суммы.
// func factorial(numbers ...int) {
// 	sum := 0
// 	for _, n := range numbers {
// 		sum += n
// 	}
// 	result := 1
// 	for i := 2; i <= sum; i++ {
// 		result *= i
// 	}
// 	fmt.Println(result)
// }

// Напишите анонимную функцию,
// которая принимает на вход слайс целых чисел и возвращает новый слайс,
// содержащий только нечетные числа из исходного списка.
// Написанную вами анонимную функцию присвойте переменной filterOddNumbers.
// func main() {
// 	filterOddNumbers := func(numbers []int) []int {
// 		newSlice := []int{}
// 		for i := range numbers {
// 			if numbers[i]%2 != 0 {
// 				newSlice = append(newSlice, numbers[i])
// 			}
// 		}
// 		return newSlice
// 	}
// 	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	fmt.Println(filterOddNumbers(mySlice))
// }

// Напишите анонимную функцию, которая принимает на вход слайс строк и функцию callback(),
// которая, в свою очередь, принимает на вход строку и возвращает булево значение.
// Функция должна вернуть новый слайс, содержащий только те строки из исходного слайса,
// для которых функция callback() возвращает true.
// Написанную вами функцию присвойте переменной filteredStrings.
// filteredStrings := func(stringSlice []string, callback func(string) bool) []string {
// 	result := []string{}
// 	for i := 0; i < len(stringSlice); i++ {
// 		if callback(stringSlice[i]) {
// 			result = append(result, stringSlice[i])
// 		}
// 	}
// 	return result
// }

// 1) создайте переменную a со значением 100
// // 2) создайте указатель *ptr на  переменную a
// // 3) присвойте значения указателя *ptr в указатель *ptr2
// // 4) присвойте значение 200 через указатель *ptr2.
// // 5) выведите содержимое обоих указателей в одной строке через пробел
// func Abc() {
// 	var a int = 100
// 	var ptr *int = &a
// 	var ptr2 *int = ptr
// 	*ptr2 = 200
// 	fmt.Println(ptr, ptr2)
// }

// func main() {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	scanner.Split(bufio.ScanWords)
// 	var myMap map[string]int = map[string]int{}
// 	for {
// 		scanner.Scan()
// 		var input string = scanner.Text()
// 		if input == "stop" {
// 			break
// 		}
// 		myMap[input] += 1
// 	}
// 	fmt.Println(myMap)
// }
// //apple banana apple pear apple banana stop
// // map[apple:3 banana:2 pear:1]
//
// func isPalindrome() {
// 	var s string
// 	fmt.Scan(&s)
// 	var runes []rune = []rune(strings.ToLower(s))
// 	length := len(runes)
// 	var ans bool = true
// 	for i := 0; i < length/2; i++ {
// 		if runes[i] != runes[length-1-i] {
// 			ans = false
// 			break
// 		}
// 	}
// 	fmt.Println(strconv.FormatBool(ans))
// }
// mapIntegers := func(nums []int, callback func(int) int) []int {
// 	var newSlice []int
// 	for i := range nums {
// 		newSlice = append(newSlice, callback(nums[i]))
// 	}
// 	return newSlice
// }
