package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	AuthorID string
	Authour  string
	TitleID  string
	Title    string
	InfoURL  string
	ZipURL   string
}

func findEntries(siteURL string) ([]Entry, error) {
	doc, err := goquery.NewDocument(siteURL)
	if err != nil {
		return nil, err
	}
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		println(elem.Text(), elem.AttrOr("href", ""))
	})
	return nil, nil
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person1779.html"
	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Title, entry.ZipURL)
	}
}
