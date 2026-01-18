package additional

// type Data struct {
// 	People []People `json:"people"`
// }
// type People struct {
// 	Id        int    `json:"id"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Email     string `json:"email"`
// 	Gender    string `json:"gender"`
// 	IpAddress string `json:"ip_address"`
// }

// func main() {
// 	file, err := os.Open("ex3file.json")
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	var data Data
// 	err = json.NewDecoder(file).Decode(&data)
// 	if err != nil {
// 		fmt.Println("Error decoder", err)
// 		return
// 	}
// 	myFile, err := os.Create("answer.txt")
// 	if err != nil {
// 		fmt.Println("error creating file", err)
// 	}
// 	defer myFile.Close()
// 	for i := range data.People {
// 		if data.People[i].Gender == "Female" {
// 			var str string = data.People[i].FirstName + " " + data.People[i].LastName + "\n"
// 			_, err := myFile.WriteString(str)
// 			if err != nil {
// 				fmt.Println("Error writing", err)
// 			}
// 		}
// 	}
// 	fmt.Println("Completed")
// }
