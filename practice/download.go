package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	// httpサーバーからのダウンロード
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV:", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Println("cannot read content:", err)
			continue
		}
		ch <- b
	}
}

func main() {
	urls := []string{
		"csvファイル名1",
		"csvファイル名2",
		"csvファイル名3",
	}
	// バイト列を転送するためのchannelを作成
	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				log.Fatal(err)
			}
			// レコードの登録
			insertRecords(records)
		}
	}
	wg.Wait()
}

func insertRecords(records []string) {
	fmt.Println(records)
}
