package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@dbserver/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	id := 3
	// ステートメントの作成
	rows, err := db.Query(`SELECT name, age FROM users WHERE id < $1`, id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// レコードをフェッチ
	for rows.Next() {
		var name string
		var age int
		//  値の抽出
		err = rows.Scan(&name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name=%v, age=%v\n", name, age)
	}
}
