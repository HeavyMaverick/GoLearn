package additional

// type Group struct {
// 	Id      int       `json:"id"`
// 	Course  int       `json:"course"`
// 	Student []Student `json:"students"`
// }

// type Student struct {
// 	FirstName string `json:"FirstName"`
// 	LastName  string `json:"LastName"`
// 	Rating    []int  `json:"Rating"`
// }

// func main() {
// 	file, err := os.Open("ex4file.json")
// 	if err != nil {
// 		fmt.Println("Open file err", err)
// 	}
// 	defer file.Close()
// 	var Group Group
// 	err = json.NewDecoder(file).Decode(&Group)
// 	if err != nil {
// 		fmt.Println("Error decoding", err)
// 	}
// 	var kolichestvo int
// 	var sumocen int
// 	var sred float64
// 	for i, v := range Group.Student {
// 		for j := 0; j < len(Group.Student[i].Rating); j++ {
// 			sumocen += v.Rating[j]
// 		}
// 		kolichestvo += len(v.Rating)
// 	}
// 	sred = float64(sumocen) / float64(kolichestvo)
// 	fmt.Printf("%.2f\n", sred)
// }
