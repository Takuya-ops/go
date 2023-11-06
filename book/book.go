package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"regexp"
	"strings"

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
	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html", token[1], token[2])
		authour, zipURL := findAuthourAndZIP(pageURL)
		println(zipURL)
		println(authour)
		// println(elem.Text(), elem.AttrOr("href", ""))

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

func findAuthourAndZIP(siteURL string) (string, string) {
	doc, err := goquery.NewNewDocument(siteURL)
	if err != nil {
		return "", ""
	}
	authour := doc.Find("table[summary=作家データ]tr:nth-child(1) td:nth-child(2)").Text()
	zipURL := ""
	doc.Find("table.downloaded a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, ".zip") {
			zipURL = href
		}
	})
	return authour, zipURL
}


func extractText(zipURL string) (string, error) {
	resp, err := http.Get(zipURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b))
	if err != nil {
		return "", err
	}
	for _, file := range r.File {
		if path.Ext(file.Name) == ".txt" {
			f, err := file.Open()
			if err != nil {
				return "", err
			}
			b, err := ioutil.ReadAll(f)
			f.Close()
			if err != nil {
				return "", err
			}
			decoder := japanese.ShiftJIS.NewDecoder()
			b, err = decoder.Bytes(b)
			if err != nil {
				return "", err
			}
			return string(b), nil
		}
	}
	return "", errors.New("コンテンツが見つかりません")
}
