package additional

import (
	"fmt"
	"io"
	"strings"
)

type Parent struct {
	Name     string
	Children []Child
}

type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}
	var abc []Child
	return Parent{
		Name:     p.Name,
		Children: append(abc, p.Children...),
	}
}

func mmm() {
	cp := CopyParent(nil) // -> Parent{}
	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}
	cp = CopyParent(p)
	// при мутациях в копии "cp"
	// изначальная структура "p" не изменяется
	cp.Children[0] = Child{
		Name: "Gosha",
		Age:  30,
	}
	fmt.Println(p.Children)  // -> [{Andy 18}]
	fmt.Println(cp.Children) // -> [{Gosha 30}]
	r := strings.NewReader("Привет")
	processData(r)
}

func processData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		n, err := reader.Read(b)
		if n == 0 {
			break
		}
		fmt.Printf("Read %v bytes: %v\n", n, string(b))
		if err == io.EOF {
			return
		}
	}
}

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
