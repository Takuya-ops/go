package main

import (
	"encoding/json"
	"fmt"
)

// structの定義
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// エンコードするデータ（Person型の変数）
	person := Person{
		Name:  "John",
		Age:   25,
		Email: "john@example.com",
	}

	// データをJSON形式にエンコード（marshal）
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// エンコードされたJSONデータを表示
	fmt.Println("Encoded JSON Data:", string(jsonData))

	// エンコードされたJSONデータをデコード（unmarshal）
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// デコードされたデータを表示
	fmt.Println("Decoded Person:")
	fmt.Println("Name:", decodedPerson.Name)
	fmt.Println("Age:", decodedPerson.Age)
	fmt.Println("Email:", decodedPerson.Email)
}
