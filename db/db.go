package main

import (
	"database/sql"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := []string{
		`CREATE TABLE IF NOT EXISTS authours(authour_id TEXT, authour TEXT, PRIMARY KEY(authour_id))`,
		`CREATE TABLE IF NOT EXISTS contents(authour_id TEXT, title_id TEXT, title TEXT, content TEXT, PRIMARY KEY(authour_id, title_id))`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS contents_fts USING fts4(words)`,
	}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	b, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err = japanese.ShiftJIS.NewDecorder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	content := string(b)
	res, err := db.Exec(`INSERT INTO contents(authour_id, title_id, title, content) values (?,?,?,?)`, "000879", "14", "サンプルです", content)
	if err != nil {
		log.Fatal(err)
	}
	docID, err := res.LastInsertId()
}
