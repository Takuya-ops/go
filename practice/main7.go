package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Gender  string
	Address string
}

func main() {
	// 1. Person構造体のインスタンスを作成し、JSONエンコード
	person := Person{
		Name:    "jon lenon",
		Age:     30,
		Gender:  "male",
		Address: "123 street",
	}

	// 2.jsonエンコードされたデータを表示
	encodedData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("エンコードエラー：", err)
		return
	}
	fmt.Println("JSONエンコード結果：", string(encodedData))

	// 3. jsonデータをデコード
	var decodedPerson Person
	if err := json.Unmarshal(encodedData, &decodedPerson); err != nil {
		fmt.Println("デコードエラー：", err)
		return
	}

	// 4. 復元されたPerson構造体を表示
	fmt.Println("デコード結果：", decodedPerson)
	fmt.Println("復元されたPerson構造体:")
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", decodedPerson.Name, decodedPerson.Age, decodedPerson.Address)
}
