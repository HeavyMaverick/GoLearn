package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type JsonData struct {
	Rates map[string]float64 `json:"rates"`
}

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://open.er-api.com/v6/latest/USD", nil)
	if err != nil {
		fmt.Println("Request error", err)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Response err", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Reading error", err)
	}
	fmt.Println(string(body))
	var JsonData JsonData
	err = json.Unmarshal(body, &JsonData)
	if err != nil {
		fmt.Println("Unmarshal err", err)
	}
	fmt.Println(JsonData.Rates)
	var counter int = 0
	for k := range JsonData.Rates {
		if strings.HasPrefix(k, "T") {
			counter += 1
		}
	}
	fmt.Println(counter)
}
