package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://api.chucknorris.io/jokes/rBOPF3bYRN6S6P2Y2ZCCWw", nil)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Header.Get("created_at"))
	ans := struct {
		CreatedAt string `json:"created_at"`
	}{}
	err = json.Unmarshal(body, &ans)
	if err != nil {
		fmt.Println("Error unmarshal", err)
	}
	fmt.Println(ans.CreatedAt)
}
