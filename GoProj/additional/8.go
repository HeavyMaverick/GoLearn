package additional

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// type Parent struct {
// 	Name     string
// 	Children []Child
// }

// type Child struct {
// 	Name string
// 	Age  int
// }

// func CopyParent(p *Parent) Parent {
// 	if p == nil {
// 		return Parent{}
// 	}
// 	var abc []Child
// 	return Parent{
// 		Name:     p.Name,
// 		Children: append(abc, p.Children...),
// 	}
// }

// func Mmm() {
// 	cp := CopyParent(nil) // -> Parent{}
// 	p := &Parent{
// 		Name: "Harry",
// 		Children: []Child{
// 			{
// 				Name: "Andy",
// 				Age:  18,
// 			},
// 		},
// 	}
// 	cp = CopyParent(p)
// 	// при мутациях в копии "cp"
// 	// изначальная структура "p" не изменяется
// 	cp.Children[0] = Child{
// 		Name: "Gosha",
// 		Age:  30,
// 	}
// 	fmt.Println(p.Children)  // -> [{Andy 18}]
// 	fmt.Println(cp.Children) // -> [{Gosha 30}]
// 	r := strings.NewReader("Привет")
// 	ProcessData(r)
// }

// func ProcessData(reader io.Reader) {
// 	b := make([]byte, 2)
// 	for {
// 		n, err := reader.Read(b)
// 		if n == 0 {
// 			break
// 		}
// 		fmt.Printf("Read %v bytes: %v\n", n, string(b))
// 		if err == io.EOF {
// 			return
// 		}
// 	}
// }

// 123abc123
// sum 123 123
// func main() {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	s := scanner.Text()
// 	var currentNumber string
// 	var sum int
// 	for _, x := range s {
// 		if unicode.IsDigit(x) {
// 			currentNumber += string(x)
// 		} else {
// 			if currentNumber != "" {
// 				num, _ := strconv.Atoi(currentNumber)
// 				sum += num
// 				currentNumber = ""
// 			}
// 		}
// 	}
// 	if currentNumber != "" {
// 		num, _ := strconv.Atoi(currentNumber)
// 		sum += num
// 	}
// 	fmt.Println(sum)
// }

// Кольцевой буфер
// 5
// Добавлен 0: [0]
// Добавлен 1: [0 1]
// Добавлен 2: [0 1 2]
// Добавлен 3: [0 1 2 3]
// Добавлен 4: [0 1 2 3 4]
// Добавлен 5: [1 2 3 4 5]
// Добавлен 6: [2 3 4 5 6]
// func main() {
// 	var capacity int
// 	_, err := fmt.Scan(&capacity)
// 	if err != nil {
// 		fmt.Println("Err", err)
// 	}
// 	queue := make([]int, 0, capacity)
// 	if capacity == 0 {
// 		queue = append(queue, 0)
// 		fmt.Printf("Добавлен 0: %v\n", queue)
// 		return
// 	}
// 	for i := range 10 {
// 		if len(queue) == capacity {
// 			queue = queue[1:]
// 		}
// 		queue = append(queue, i)
// 		fmt.Printf("Добавлен %d: %v\n", i, queue)
// 	}
// }

// 1 2 3 4
// 3 4 5 6
// Sample Output:
// [1 2 3 4 5 6]
// func union(a, b []int) []int {
// 	result := make([]int, 0)
// 	seen := map[int]bool{}
// 	for _, x := range a {
// 		if !seen[x] {
// 			seen[x] = true
// 			result = append(result, x)
// 		}
// 	}
// 	for _, y := range b {
// 		if !seen[y] {
// 			seen[y] = true
// 			result = append(result, y)
// 		}
// 	}
// 	return result
// }
//
//
// 1
// 2
// 3
// ----
// 6
// func main() {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	var j int
// 	for scanner.Scan() {
// 		if scanner.Text() == "" {
// 			break
// 		}
// 		num, err := strconv.Atoi(scanner.Text())
// 		if err != nil {
// 			continue
// 		}
// 		j += num
// 	}
// 	_, err := io.WriteString(os.Stdout, strconv.Itoa(j))
// 	if err != nil {
// 		println("Error", err)
// 	}
// }

// Print word w/max length in str
// func main() {
// 	sc := bufio.NewScanner(os.Stdin)
// 	sc.Scan()
// 	str := sc.Text()
// 	var nStr string
// 	var strSlc []string
// 	for _, runee := range str {
// 		if unicode.IsLetter(runee) && string(runee) != " " {
// 			nStr += string(runee)
// 		} else {
// 			if nStr != "" {
// 				strSlc = append(strSlc, nStr)
// 				nStr = ""
// 			}
// 		}
// 	}
// 	var maxString string
// 	var maxlen int
// 	for i, x := range strSlc {
// 		var n int = utf8.RuneCountInString(strSlc[i])
// 		if n >= maxlen {
// 			maxlen = n
// 			maxString = x
// 		}
// 	}
// 	fmt.Println(maxString)
// }

// 2007-01-01
// 2007-01-06
// 5
// func main() {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	var strTime1, strTime2 string
// 	scanner.Scan()
// 	strTime1 = scanner.Text()
// 	scanner.Scan()
// 	strTime2 = scanner.Text()
// 	layout := "2006-01-02"
// 	t1, err := time.Parse(layout, strTime1)
// 	if err != nil {
// 		fmt.Println("Введенные значения некорректны!")
// 		return
// 	}
// 	t2, err := time.Parse(layout, strTime2)
// 	if err != nil {
// 		fmt.Println("Введенные значения некорректны!")
// 		return
// 	}
// 	if t2.After(t1) {
// 		diff := t2.Sub(t1)
// 		hours := int(diff.Hours())
// 		fmt.Println(hours / 24)
// 	} else {
// 		if t1.After(t2) {
// 			diff := t1.Sub(t2)
// 			hours := int(diff.Hours())
// 			fmt.Println(hours / 24)
// 		} else {
// 			fmt.Println("Error")
// 		}
// 	}
// }
//
// Hello, World!
// HelloWorld
// func main() {
// 	var scanner = bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	inputString := scanner.Text()
// 	var newString strings.Builder
// 	for _, runee := range inputString {
// 		if (unicode.IsLetter(runee) || unicode.IsDigit(runee)) && string(runee) != " " {
// 			newString.WriteString(string(runee))
// 		}
// 	}
// 	fmt.Println(newString.String())
// }
// John 30 street
// {"name":"John","age":30,"address":"street"}
// func main() {
// 	var name, address string
// 	var age int
// 	_, err := fmt.Scanf("%s %d %s", &name, &age, &address)
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	var p Person = Person{
// 		Name:    name,
// 		Age:     age,
// 		Address: address,
// 	}
// 	jsonBytes, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	fmt.Printf("%s", jsonBytes)
// }

// sum all nums in file
// exmple : 123asdasd12asdasd
// output: 135
// func main() {
// 	bytes, err := os.ReadFile("ex1file.text")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	var currentNumber string
// 	var sum int
// 	for _, v := range bytes {
// 		if unicode.IsDigit(rune(v)) {
// 			currentNumber += string(v)
// 		} else {
// 			if currentNumber != "" {
// 				n, _ := strconv.Atoi(currentNumber)
// 				sum += n
// 				currentNumber = ""
// 			}
// 		}
// 	}
// 	fmt.Println(sum)
// }
// Sample Input:
// AaaAaaA
// Sample Output:
// aaaa
// func main() {
// 	bytes, err := os.ReadFile("ex2file.text")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	file, err := os.Create("newFile.txt")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	defer file.Close()
// 	var bb []byte
// 	for _, value := range bytes {
// 		if unicode.IsLower(rune(value)) {
// 			bb = append(bb, value)
// 		}
// 	}
// 	_, err = file.Write(bb)
// 	if err != nil {
// 		fmt.Println("Error")
// 	}
// 	fmt.Println("Completed")
// }

// in
// aaabbbc
// out
// a3b3c1
// func main() {
// 	var myString string
// 	_, err := fmt.Scanf("%s", &myString)
// 	if err != nil {
// 		fmt.Println("error", err)
// 	}
// 	if len(myString) == 0 {
// 		fmt.Println("")
// 		return
// 	}
// 	var count int = 1
// 	var current byte = myString[0]
// 	var result strings.Builder
// 	for i := 1; i < len(myString); i++ {
// 		if myString[i] == current {
// 			count++
// 		} else {
// 			result.WriteByte(current)
// 			result.WriteString(strconv.Itoa(count))
// 			current = myString[i]
// 			count = 1
// 		}
// 	}
// 	result.WriteByte(current)
// 	result.WriteString(strconv.Itoa(count))
// 	fmt.Println(result.String())
// }

// ---------------------------
// Напишите программу,
// которая принимает строку в формате json и распаковывает ее в структуру Person,
// переменную в которой будет храниться структура назовите p.
// Если формат введенной строки не соответствует формату json
// выведите  "Некорректное значение json!" и завершите выполнение программы.
// type Person struct {
// 	Name    string `json:"name"`
// 	Age     int    `json:"age"`
// 	Address string `json:"address"`
// }

// func main() {
// 	var stringJson string
// 	_, err := fmt.Scanln(&stringJson)
// 	if err != nil {
// 		fmt.Println("Некорректное значение json!")
// 	}
// 	var p Person = Person{}
// 	err = json.Unmarshal([]byte(stringJson), &p)
// 	if err != nil {
// 		fmt.Println("Некорректное значение json!")
// 		return
// 	}
// 	fmt.Println(p)
// }
